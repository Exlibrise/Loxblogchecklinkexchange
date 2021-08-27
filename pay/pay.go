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

// Client is an HTTP client, which handles "Payment Required" interruptions transparently.
// It must be initially set up with a connection the Lightning Network node that should handle the payments
// and from then on it's meant to be used as an alternative t