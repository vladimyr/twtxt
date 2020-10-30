package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/andyleap/microformats"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type HApp struct {
	Name string
	URL  *url.URL
	Logo *url.URL
}

func (h HApp) String() string {
	if h.Name == "" {
		return h.URL.Hostname()
	}

	return h.Name
}

func GetIndieClientInfo(conf *Config, client_id string) (h HApp, err error) {
	u, err := url.Parse(client_id)
	if err != nil {
		log.WithError(err).Errorf("error parsing  url: %s", client_id)
		return h, err
	}
	h.URL = u

	res, err := Request(conf, "GET", client_id, nil)
	if err != nil {
		log.WithError(err).Errorf("error making client request to %s", client_id)
		return h, err
	}
	defer res.Body.Close()

	body, err := html.Parse(res.Body)
	if err != nil {
		log.WithError(err).Errorf("error parsing source %s", client_id)
		return h, err
	}

	p := microformats.New()
	data := p.ParseNode(body, u)

	h.URL = u

	getHApp := func(data *microformats.Data) (*microformats.MicroFormat, error) {
		if data != nil {
			for _, item := range data.Items {
				if HasString(item.Type, "h-app") {
					return item, nil
				}
			}
		}
		return nil, errors.New("error: no entry found")
	}

	happ, err := getHApp(data)
	if err != nil {
		return h, err
	}

	if names, ok := happ.Properties["name"]; ok && len(names) > 0 {
		if name, ok := names[0].(string); ok {
			h.Name = name
		}
	}

	if logos, ok := happ.Properties["logo"]; ok && len(logos) > 0 {
		if logo, ok := logos[0].(string); ok {
			if u, err := url.Parse(logo); err != nil {
				h.Logo = u
			} else {
				log.WithError(err).Warnf("error parsing logo %s", logo)
			}
		}
	}

	return h, nil
}

func ValidateIndieRedirectURL(client_id, redirect_url string) error {
	u1, err := url.Parse(client_id)
	if err != nil {
		log.WithError(err).Errorf("error parsing  url: %s", client_id)
		return err
	}

	u2, err := url.Parse(client_id)
	if err != nil {
		log.WithError(err).Errorf("error parsing  url: %s", redirect_url)
		return err
	}

	if u1.Scheme != u2.Scheme {
		return errors.New("invalid redirect url, mismatched scheme")
	}

	if u1.Hostname() != u2.Hostname() {
		return errors.New("invalid redirect url, mismatched hostname")
	}

	if u1.Port() != u2.Port() {
		return errors.New("invalid redirect url, mismatched port")
	}

	return nil
}

// IndieAuthHandler ...
func (s *Server) IndieAuthHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := NewContext(s.config, s.db, r)

		nick := NormalizeUsername(p.ByName("nick"))
		if nick == "" {
			if r.Method == http.MethodHead {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			ctx.Error = true
			ctx.Message = "No user specified"
			s.render("error", w, ctx)
			return
		}

		me := r.FormValue("me")
		client_id := r.FormValue("client_id")
		redirect_url := r.FormValue("redirect_uri")
		state := r.FormValue("state")

		if me == "" || client_id == "" || redirect_url == "" || state == "" {
			log.Warn("missing authentication parameters")

			if r.Method == http.MethodHead {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			ctx.Error = true
			ctx.Message = "Error one or more authentication parameters are missing"
			s.render("error", w, ctx)
			return
		}

		response_type := r.FormValue("response_type")
		if response_type == "" {
			response_type = "id"
		}
		response_type = strings.ToLower(response_type)

		user, err := s.db.GetUser(nick)
		if err != nil {
			log.WithError(err).Errorf("error loading feed object for %s", nick)

			if r.Method == http.MethodHead {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			ctx.Error = true
			ctx.Message = "Error loading user"
			s.render("error", w, ctx)
			return
		}

		happ, err := GetIndieClientInfo(s.config, client_id)
		if err != nil {
			log.WithError(err).Warnf("error retrieving client information from %s", client_id)
		}

		if err := ValidateIndieRedirectURL(client_id, redirect_url); err != nil {
			log.WithError(err).Errorf("error validating redirect_url %s from client %s", redirect_url, client_id)

			if r.Method == http.MethodHead {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			ctx.Error = true
			ctx.Message = "Error validating redirect url"
			s.render("error", w, ctx)
			return
		}

		log.Infof("IndieAuth Authorization request for %s from %s", user.Username, happ)

		ctx.Title = "Authorize Login"
		ctx.Message = fmt.Sprintf(
			"The app %s wants to access your account %s",
			happ, user.Username,
		)
		ctx.Callback = fmt.Sprintf(
			"%s/indieauth/callback?redirect_url=%s&state=%s",
			UserURL(user.URL), redirect_url, state,
		)
		log.Infof("Callback: %q", ctx.Callback)
		s.render("prompt", w, ctx)
	}
}

// IndieAuthCallbackHandler ...
func (s *Server) IndieAuthCallbackHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		redirect_url := r.FormValue("redirect_url")
		state := r.FormValue("state")

		if redirect_url == "" || state == "" {
			log.Warn("missing callback parameters")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		u, err := url.Parse(redirect_url)
		if err != nil {
			log.WithError(err).Errorf("error parsing redirect url %s", redirect_url)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		code := GenerateToken()

		v := url.Values{}
		v.Add("code", code)
		v.Add("state", state)
		u.RawQuery = v.Encode()

		log.Infof("u.String(): %q", u.String())

		http.Redirect(w, r, u.String(), http.StatusFound)
	}
}

// IndieAuthVerifyHandler ...
func (s *Server) IndieAuthVerifyHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		nick := NormalizeUsername(p.ByName("nick"))
		if nick == "" {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		client_id := r.FormValue("client_id")
		redirect_uri := r.FormValue("redirect_uri")
		code := r.FormValue("code")

		if client_id == "" || redirect_uri == "" || code == "" {
			log.WithField("client_id", client_id).
				WithField("redirect_uri", redirect_uri).
				WithField("code", code).
				Warn("missing verification parameters")
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// TODO: Validate code was issued for client_id and redirect_uri
		if tokenCache.Get(code) == 0 {
			log.Warnf("invalid code %s for indieauth verification", code)
			// TODO: Return JSON error code: {"error": "invalid_grant"}
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		me := map[string]string{
			"me": URLForUser(s.config, nick),
		}

		data, err := json.Marshal(me)
		if err != nil {
			log.WithError(err).Error("error serializing me response")
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
