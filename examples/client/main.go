package main

import (
	"fmt"
	"io/ioutil"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/pay"
)

func main() {
	// Set up client
	lndOptions := ln.LNDoptions{ // Default 