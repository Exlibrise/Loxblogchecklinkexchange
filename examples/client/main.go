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
		MacaroonFile: "admin.macaroon", // admin.macaroon is required for making payments
	}
	lnClient, err := ln.NewLNDclient(lndOptions)
	if err != nil {
		panic(err)
	}
	client := pay.NewClient(nil, lnClient) // Uses http.DefaultClient if no http.Client is passed

	// Send request to an ln-paywalled API
	res, err := client.Get("http://localhost:8080/ping")
	if err != nil {