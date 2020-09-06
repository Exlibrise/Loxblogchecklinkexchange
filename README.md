
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

How it works
------------

With `ln-paywall` you can simply use one of the provided middlewares in your Go web service to have your web service do two things:

1. The first request gets rejected with the `402 Payment Required` HTTP status, a `Content-Type: application/vnd.lightning.bolt11` header and a Lightning ([BOLT-11](https://github.com/lightningnetwork/lightning-rfc/blob/master/11-payment-encoding.md)-conforming) invoice in the body
2. The second request must contain a `X-Preimage` header with the preimage of the paid Lightning invoice (hex encoded). The middleware checks if 1) the invoice was paid and 2) not already used for a previous request. If both preconditions are met, it continues to the next middleware or final request handler.

Prerequisites
-------------

There are currently two prerequisites:

1. A running Lightning Network node. The middleware connects to the node for example to create invoices for a request. The `ln` package currently provides factory functions for the following LN implementations:
	- [X] [lnd](https://github.com/lightningnetwork/lnd)
		- Requires the node to listen to gRPC connections
		- If you don't run it locally, it needs to listen to connections from external machines (so for example on 0.0.0.0 instead of localhost) and has the TLS certificate configured to include the external IP address of the node.
	- [X] [c-lightning](https://github.com/ElementsProject/lightning) with [Lightning Charge](https://github.com/ElementsProject/lightning-charge)
		- Run for example with Docker: ``docker run -d -u `id -u` -v `pwd`/data:/data -p 9112:9112 -e API_TOKEN=secret shesek/lightning-charge``
		- Vanilla c-lightning (without Lightning Charge) won't be supported as long as c-lightning's RPC API only works via Unix socket and cannot be used as a remote server, because this is not a good fit for potentially multiple web service instances elastically scaled across a cluster of host machines
	- [ ] [eclair](https://github.com/ACINQ/eclair) (not implemented yet - [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) )
	- Roll your own!
		- Just implement the simple `wall.LNClient` interface (only two methods!)
2. A supported storage mechanism. It's used to cache preimages that have been used as a payment for an API call, so that a user can't do multiple requests with the same preimage of a settled Lightning payment. The `wall` package currently provides factory functions for the following storages:
	- [X] A simple Go map
		- The fastest option, but 1) can't be used across horizontally scaled service instances and 2) doesn't persist data, so when you restart your server, users can re-use old preimages
	- [X] [bbolt](https://github.com/coreos/bbolt) - a fork of [Bolt](https://github.com/boltdb/bolt) maintained by CoreOS
		- Very fast, doesn't require any remote or local TCP connections and persists the data, but can't be used across horizontally scaled service instances because it's file-based. Production-ready for single-instance web services though.
	- [X] [Redis](https://redis.io/)
		- Although the slowest of these options, still fast and most suited for popular web services: Requires a remote or local TCP connection and some administration, but allows data persistency and can even be used with a horizontally scaled web service
		- Run for example with Docker: `docker run -d -p 6379:6379 redis`
			- Note: In production you should use a configuration with password (check out [`bitnami/redis`](https://hub.docker.com/r/bitnami/redis/) which makes that easy)!
	- [ ] [groupcache](https://github.com/golang/groupcache) (not implemented yet - [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) )
	- Roll your own!
		- Just implement the simple `wall.StorageClient` interface (only two methods!)

Usage
-----

[![GoDoc](http://www.godoc.org/github.com/philippgille/ln-paywall?status.svg)](http://www.godoc.org/github.com/philippgille/ln-paywall)

Get the package with `go get -u github.com/philippgille/ln-paywall/...`.

We strongly encourage you to use vendoring, because as long as `ln-paywall` is version `0.x`, breaking changes may be introduced in new versions, including changes to the package name / import path. The project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html) and all notable changes to this project are documented in [RELEASES.md](https://github.com/philippgille/ln-paywall/blob/master/RELEASES.md).

### Middleware

The best way to see how to use `ln-paywall` is by example. In the below examples we create a web service that responds to requests to `/ping` with "pong", using Gin as the web framework.

```Go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
)

func main() {
	r := gin.Default()

	// Configure middleware
	invoiceOptions := wall.DefaultInvoiceOptions // Price: 1 Satoshi; Memo: "API call"
	lndOptions := ln.DefaultLNDoptions           // Address: "localhost:10009", CertFile: "tls.cert", MacaroonFile: "invoice.macaroon"
	storageClient := storage.NewGoMap()          // Local in-memory cache
	lnClient, err := ln.NewLNDclient(lndOptions)
	if err != nil {
		panic(err)
	}
	// Use middleware
	r.Use(wall.NewGinMiddleware(invoiceOptions, lnClient, storageClient))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.Run() // Listen and serve on 0.0.0.0:8080
}
```

This is just the most basic example. See the list of examples below for examples with other web frameworks / routers / just the stdlib, as well as for a more complex and useful example.

#### List of examples

Follow the links to the example code files.

Simple examples to show the use for the different web frameworks / routers / just the stdlib:

- [Gin](examples/ping/gin/main.go)
- [Gin (with c-lightning + Lightning Charge as backend)](examples/ping/gin-charge/main.go)
- [gorilla/mux](examples/ping/gorilla-mux/main.go)
- [net/http HandlerFunc](examples/ping/handlerfunc/main.go)
- [Echo](examples/ping/echo/main.go)

More complex and useful example:

- [QR code generation API using Gin](examples/qr-code/main.go)
	- Ready-to-use Docker image: [https://hub.docker.com/r/philippgille/qr-code/](https://hub.docker.com/r/philippgille/qr-code/)
	- Try out the demo deployed on [https://lightning.ws](https://lightning.ws)

### Client

```Go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/pay"
)

func main() {
	// Set up client