
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
	- [Middleware](#middleware)
		- [List of examples](#list-of-examples)
	- [Client](#client)
- [Related projects](#related-projects)

Purpose
-------

Until the rise of cryptocurrencies, if you wanted to monetize your API (set up a paywall), you had to:

1. Use a centralized service (like PayPal)
    - Can shut you down any time
    - High fees
    - Your API users need an account
    - Can be hacked
2. Keep track of your API users (keep accounts and their API keys in some database)
    - Privacy concerns
    - Data breaches / leaks
3. Charge for a bunch of requests, like 10.000 at a time, because real per-request payments weren't possible

With cryptocurrencies in general some of those problems were solved, but with long confirmation times and high per-transaction fees a real per-request billing was still not feasable.

But then came the [Lightning Network](https://lightning.network/), an implementation of routed payment channels, which enables *real* **near-instant microtransactions** with **extremely low fees**, which cryptocurrencies have long promised, but never delivered. It's a *second layer* on top of existing cryptocurrencies like Bitcoin that scales far beyond the limitations of the underlying blockchain.

`ln-paywall` makes it easy to set up an API paywall for payments over the Lightning Network.
