package pay

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// LNclient is the abstraction of a Lightning Network node client for paying LN invoices.
type LNclient interface {
	// Pay pays the invoice and returns the preimage on success, or an error on failure.
	Pay(invoice string) (string, error)
}

// Client is an HTTP client, which handles "Payment Required" interru