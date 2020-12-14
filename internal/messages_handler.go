package internal

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"

	"github.com/jointwt/twtxt/types"
)

const (
	mailboxesDir = "mailboxes"
)

// MailboxHandler ...
func (s *Server) MailboxHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := NewContext(s.config, s.db, r)

		ctx.Error = true
		ctx.Message = "Not Implemented Yet"
		s.render("error", w, ctx)
		return
	}
}

// MessagesHandler ...
func (s *Server) MessagesHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := NewContext(s.config, s.db, r)

		twts := s.cache.GetBySuffix(fmt.Sprintf(":%s", ctx.User.Username), false)
		sort.Sort(sort.Reverse(twts))

		var pagedTwts types.Twts

		page := SafeParseInt(r.FormValue("p"), 1)
		pager := paginator.New(adapter.NewSliceAdapter(twts), s.config.TwtsPerPage)
		pager.SetPage(page)

		if err := pager.Results(&pagedTwts); err != nil {
			ctx.Error = true
			ctx.Message = "An error occurred while loading search results"
			s.render("error", w, ctx)
			return
		}

		if r.Method == http.MethodHead {
			defer r.Body.Close()
			return
		}

		ctx.Title = "Private Messages"

		ctx.Twts = FilterTwts(ctx.User, pagedTwts)
		ctx.Pager = &pager
		s.render("messages", w, ctx)
		return
	}
}
