/*
Package wall contains all paywall-related code.

This is the package you need to use for creating a middleware for one of the supported handlers / routers / frameworks.
For creating a middleware you only need to call one of the provided factory functions,
but all functions require a storage client (an implementation of the wall.StorageClient interface) as parameter.
You can either pick one from the storage package (https://www.godoc.org/github.com/philippgille/ln-paywall/storage), or implement your own.

Usage

Here's one example of a web service implemented with Gin.
For more examples check out the "examples" directory in the GitHub repository of this package (https://github.com/philippgille/ln-paywall/tree/master/examples).

	package main

	import (
		"net/http"

		"github.com/gin-gonic/gin"
		"github.com/philippgille/ln-paywall/ln"
		"github.co