
ln-paywall
==========

[![GoDoc](http://www.godoc.org/github.com/philippgille/ln-paywall/wall?status.svg)](http://www.godoc.org/github.com/philippgille/ln-paywall/wall) [![Build Status](https://travis-ci.org/philippgille/ln-paywall.svg?branch=master)](https://travis-ci.org/philippgille/ln-paywall) [![Go Report Card](https://goreportcard.com/badge/github.com/philippgille/ln-paywall)](https://goreportcard.com/report/github.com/philippgille/ln-paywall) [![GitHub Releases](https://img.shields.io/github/release/philippgille/ln-paywall.svg)](https://github.com/philippgille/ln-paywall/releases)

Go middleware for monetizing your API on a per-request basis with Bitcoin and Lightning ⚡️

Middlewares for:

- [X] [net/http](https://golang.org/pkg/net/http/) `HandlerFunc`
- [X] [net/http](https://golang.org/pkg/net/http/) `Handler`
	- Compatible with routers like [gorilla/mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter) and [chi](https://github.com/go-chi/chi)
- [X] [Gin](https://github.com/gin-gonic/gin)
- [X] [Echo](https://github.com/labstack/echo)

A client package exists as well to make *consuming* LN-paywalled APIs extremely easy (you just use it like the standard Go `http.Client` and the payment handling is done in the background).

An API gateway is on the roadmap as well, which you can use to monetize your API that's written in *any* language, not just in Go.

Contents
--------

- [Purpose](#purpose)
- [How it works](#how-it-works)
- [Prerequisites](#prerequisites)
- [Usage](#usage)