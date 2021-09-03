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
// and from then on it's meant to be used as an alternative to the "net/http.Client".
// The calling code only needs to call the Do(...) method once, instead of handling
// "402 Payment Required" responses and re-sending the original request after payment.
type Client struct {
	c *http.Client
	l LNclient
}

// Get sends an HTTP GE