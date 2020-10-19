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
	lndOptions := ln.DefaultLNDoptions           // Address: "localhos