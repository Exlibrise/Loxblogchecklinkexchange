package ln

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
)

// stdOutLogger logs to stdout, while the default log package loggers log to stderr.
var stdOutLogger = log.New(os.Stdout, "", log.LstdFlags)

// Invoice is a Lightning Network invoice and contains the typical invoice string and the payment hash.
type Invoice struct {
	// The unique identifier for the invoice in the LN node.
	// The value depends on the LN node implementation.
	//
	// For example, lnd uses the payment hash (a.k.a. preimage hash) as ID.
	// It doesn't use this term ("ID"), but when fetching a single invoice via RPC,
	// the payment hash is used as identifier.
	// Also, Lightning Lab's (creators of lnd) desktop app "Lightning" explicitly shows
	// the payment hash in a field with the title "invoice ID" in the