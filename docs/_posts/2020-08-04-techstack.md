---
layout: page
title: "Tech Stack"
category: doc
date: 2020-08-04 14:38:00
order: 3
---

Currently twtxt consists of the following components:

- [Go](https://golang.org) -- A Go backend mostly written from scratch from
  bits and pieces borrowed from other projects. It uses primarily the
  [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
  mux library for routing of requests for both the frontend's backend and
  the API.
- [Bitcask](https://github.com/prologic/bitcask) -- This is the primary
  KV store that backs most of the "metadata" of [twtxt.net](https://twtxt.net).
  It is a high-performance KV store designed for fast O(1) lookups of key/vaue
  pairs of data with 1 Disk IOPS per key.
- [PicoCSS](https://picocss.com) -- This is a "classless" CSS library that
  was chosen because of its very "lightweight" size and was quick to initially
  develop in as the primary author ([James Mills](https://github.com/prologic))
  is no frontend dev 🤣 The forked version (_only slightly modified_) of this
  library can be found [here](https://github.com/prologic/picocss)
- [UmbrellaJS](https://umbrellajs.com/) -- Again this was chosen over more
  common choices like jQuery becuase of its "lightweight" size. THe project uses
  a little JS very sparingly and has a requirement of graceful degradation so
  that users of the platform are able to use [twtxt.net](https://twtxt.net) on
  poor Internet connections or even from browsers without Javascript.

That's pretty much it for now. If you are interested in contributing in any way
we welcome you! You can find all the sources at:

- `./internal/*.go` and in some sub-directories that are sub-packages of the project.

In addition you can find all the static assets and templates here:

- `./internal/static/{css,js}`
- `./internal/templates/*.html`
