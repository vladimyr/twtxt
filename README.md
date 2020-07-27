# twtxt

![](https://github.com/prologic/twtxt/workflows/Coverage/badge.svg)
![](https://github.com/prologic/twtxt/workflows/Docker/badge.svg)
![](https://github.com/prologic/twtxt/workflows/Go/badge.svg)
![](https://github.com/prologic/twtxt/workflows/ReviewDog/badge.svg)

[![CodeCov](https://codecov.io/gh/prologic/twtxt/branch/master/graph/badge.svg)](https://codecov.io/gh/prologic/twtxt)
[![Go Report Card](https://goreportcard.com/badge/prologic/twtxt)](https://goreportcard.com/report/prologic/twtxt)
[![codebeat badge](https://codebeat.co/badges/15fba8a5-3044-4f40-936f-9e0f5d5d1fd9)](https://codebeat.co/projects/github-com-prologic-twtxt-master)
[![GoDoc](https://godoc.org/github.com/prologic/twtxt?status.svg)](https://godoc.org/github.com/prologic/twtxt) 
[![GitHub license](https://img.shields.io/github/license/prologic/twtxt.svg)](https://github.com/prologic/twtxt)
[![Sourcegraph](https://sourcegraph.com/github.com/prologic/twtxt/-/badge.svg)](https://sourcegraph.com/github.com/prologic/twtxt?badge)
[![TODOs](https://img.shields.io/endpoint?url=https://api.tickgit.com/badge?repo=github.com/prologic/twtxt)](https://www.tickgit.com/browse?repo=github.com/prologic/twtxt)

📕 twtxt is a Self-Hosted, Twitter™-like Decentralised microBlogging platform. No ads, no tracking, your content!

Technically `twtxt` is a [twtxt](https://twtxt.readthedocs.io/en/latest/) client in the form
of a web application. It supports multiple users and
also hosts user feeds directly.

There is also a publicly (_free_) service online available at:

- https://twtxt.net/

> __NOTE:__ I [James Mills](https://github.com/prologic) run this first (_of which I hope to be many_) `twtxt` instance on pretty cheap hardware on a limited budget. Please use it fairly so everyone can enjoy using it equally! Please be sure to read the [/privacy](https://twtxt.net/privacy) before singing up (_pretty striaght forward_) and happy Twt'ing! 🤗

![Screenshot 1](./screenshot1.png)
![Screenshot 2](./screenshot2.png)

## Installation

### Pre-built Binaries

As a first point, please try to use one of the pre-built binaries  that are
available on the [Releases](https://github.com/prologic/twtxt/releases) page.

### Building from source

This is an option if you are familiar with [Go](https://golang.org) development.

1. Clone this repository (_this is important_)

```console
git clone https://github.com/prologic/twtxt.git
```

2. Install required dependencies (_this is important_)

```console
make deps
```

3. Build the binaries

```console
make
```

__NOTE___: It is important you follow these steps and don't just simply attempt
           `go get ...` this project as that will not work (#30) due to the
           need to package templates and static assets which we use the
           [go.rice](https://github.com/GeertJohan/go.rice) tool for.

## Usage

### Deploy with Docker Compose

Run the compose configuration:

```console
docker-compose up -d
```

Then visit: http://localhost:8000/

### Web App

Run twtd:

```console
twtd -r
```

__NOTE:__ Registrations are disabled by default so hence the `-r` flag above.

Then visit: http://localhost:8000/

You can configure other options by specifying them on the command-line:

```console
./twtd -h
Usage of ./twtd:
  -u, --base-url string           base url to use (default "http://0.0.0.0:8000")
  -b, --bind string               [int]:<port> to bind to (default "0.0.0.0:8000")
  -S, --cookie-secret string      cookie secret to use (default "PLEASE_CHANGE_ME!!!")
  -d, --data string               data directory (default "./data")
  -D, --debug                     enable debug logging
  -F, --feed-sources strings      external feed sources (default [https://feeds.twtxt.net/we-are-feeds.txt,https://raw.githubusercontent.com/mdom/we-are-twtxt/master/we-are-bots.txt])
  -L, --max-tweet-length int      maximum length of posts (default 288)
  -n, --name string               set the instance's name (default "twtxt.net")
  -r, --register                  enable user registration
  -E, --session-expiry duration   session expiry to use (default 24h0m0s)
  -s, --store string              store to use (default "bitcask://twtxt.db")
  -t, --theme string              set the default theme (default "dark")
  -T, --tweets-per-page int       tweets per page to display (default 50)
  -v, --version                   display version information
pflag: help requested
```

## Production Deployments

### Docker Swarm

You can deploy `twtxt` to a [Docker Swarm](https://docs.docker.com/engine/swarm/)
cluster by utilsing the provided `twtxt.yaml` Docker Stack. This also depends on
and uses the [Traefik](https://docs.traefik.io/) ingres load balancer so you must
also have that configured and running in your cluster appropriately.

```console
docker stack deploy -c twtxt.yml
```

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/prologic/twtxt.svg)](https://starcharts.herokuapp.com/prologic/twtxt)

## Support

Support the ongoing development of twtxt!

**Sponser**

- Become a [Sponsor](https://www.patreon.com/prologic)
- Contribute! See [TODO](/TODO.md)

## Contributors

Thank you to all those that have contributed to this project, battle-tested it, used it in their own projects or products, fixed bugs, improved performance and even fix tiny typos in documentation! Thank you and keep contributing!

You can find an [AUTHORS](/AUTHORS) file where we keep a list of contributors to the project. If you contriibute a PR please consider adding your name there. There is also Github's own [Contributors](https://github.com/prologic/twtxt/graphs/contributors) statistics.

[![](https://sourcerer.io/fame/prologic/prologic/twtxt/images/0)](https://sourcerer.io/fame/prologic/prologic/twtxt/links/0)
[![](https://sourcerer.io/fame/prologic/prologic/twtxt/images/1)](https://sourcerer.io/fame/prologic/prologic/twtxt/links/1)

## Related Projects

- [rss2twtxt](https://github.com/prologic/rss2twtxt)

## License

twtwt is licensed under the terms of the [MIT License](/LICENSE)
