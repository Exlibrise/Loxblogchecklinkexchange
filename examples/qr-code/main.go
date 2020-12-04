package main

import (
	"flag"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
	qrcode "github.com/skip2/go-qrcode"
)

var lndAddress = flag.String("addr", "localhost:10009", "Address of the lnd node (including gRPC port)")
var dataDir = flag.String("dataDir", "data/", "Relative path to the data directory, where tls.cert and invoice.macaroon are located")
var price = flag.Int64("price", 1000, "Price of one request in Satoshis (at an exchange rate of $1,000 for 1 BTC 100