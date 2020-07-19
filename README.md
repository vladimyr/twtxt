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

twtxt is a [twtxt](https://twtxt.readthedocs.io/en/latest/) client in the form
of a web application. ~and command-line client. It supports multiple users and
also hosts user feeds directly. It also  has a builtin registry and search.~

There is also a publically (_free_) service online available at:

- https://twtxt.net/

![Screenshot 1](./screenshot1.png)
![Screenshot 2](./screenshot2.png)

## Installation

### Source

```#!bash
$ go get -u github.com/prologic/twtxt/...
```

## Usage

### CLI

Run twt:

```#!bash
$ twt
```

### Web App

Run twtd:

```#!bash
$ twtd
```

Then visit: http://localhost:8000/

## License

twtwt is licensed under the terms of the [MIT License](/LICENSE)
