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
var dataDir = flag.String("dat