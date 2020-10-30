package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/andyleap/microformats"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"

	"github.com/prologic/twtxt/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <url>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	client_id := os.Args[1]

	u, err := url.Parse(client_id)
	if err != nil {
		log.WithError(err).Fatalf("error parsing  url: %s", client_id)
	}

	conf := &internal.Config{}

	res, err := internal.Request(conf, "GET", client_id, nil)
	if err != nil {
		log.WithError(err).Fatal("error making client request")
	}
	defer res.Body.Close()

	body, err := html.Parse(res.Body)
	if err != nil {
		log.WithError(err).Fatalf("error parsing source %s", client_id)
	}

	p := microformats.New()
	data := p.ParseNode(body, u)

	out, err := json.Marshal(data)
	if err != nil {
		log.WithError(err).Fatal("error marshalling json")
	}

	fmt.Println(string(out))
}
