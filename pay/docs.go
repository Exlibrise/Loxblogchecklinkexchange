/*
Package pay contains the client-side functionality for easy use of APIs that are paywalled with ln-paywall or other compatible paywall implementations.

Usage

	package main

	import (
		"fmt"
		"io/ioutil"

		"github.com/philippgille/ln-paywall/ln"
		"github.com/philippgille/ln-paywall/pay"
	)

	func main() {
		// Set up client
		lndOptions := ln.LNDopt