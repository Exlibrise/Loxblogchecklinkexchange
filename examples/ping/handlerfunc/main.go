
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

func main() {
	// Configure middleware
	invoiceOptions := wall.DefaultInvoiceOptions // Price: 1 Satoshi; Memo: "API call"
	lndOptions := ln.DefaultLNDoptions           // Address: "localhost:10009", CertFile: "tls.cert", MacaroonFile: "invoice.macaroon"
	storageClient := storage.NewGoMap()          // Local in-memory cache
	lnClient, err := ln.NewLNDclient(lndOptions)
	if err != nil {
		panic(err)
	}