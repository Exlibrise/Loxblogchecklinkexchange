package main

import (
	"fmt"
	"io/ioutil"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/pay"
)

func main() {
	// Set up client
	lndOptions := ln.LNDoptions{ // Default address: "localhost:10009", CertFile: "tls.cert"
		MacaroonFile: "admin.macaroon", // admi