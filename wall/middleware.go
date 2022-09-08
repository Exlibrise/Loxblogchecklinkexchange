package wall

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/philippgille/ln-paywall/ln"
)

// stdOutLogger logs to stdout, while the default log package loggers log to stderr.
var stdOutLogger = log.New(os.Stdout, "", log.LstdFlags)

// InvoiceOptions are the options for an invoice.
type InvoiceOptions struct {
	// Amount of Satoshis you want to have paid for one API call.
	// Values below 1 are automatically changed to the default value.
	// Optional (1 by default).
	Price int64
	// Note to be shown on the invoice,
	// for example: "API call to api.example.com".
	// Optional ("" by default).
	Memo string
}

// DefaultInvoiceOptions provides default values for InvoiceOptions.
var DefaultInvoiceOptions = InvoiceOptions{
	Price: 1,
	Memo:  "API call",
}

// StorageClient is an abstraction for different storage client implementations.
// A storage client must be able to store and retrieve invoiceMetaData objects.
type StorageClient interface {
	// Set stores the given invoiceMetaData for the given preimage hash.
	Set(string, interface{}) error
	// Get retrieves the invoiceMetaData for the given preimage hash
	// and populates the fields of the object that the passed pointer
	// points to with the values of the retrieved object's values.
	// If no object is found it returns (false, nil).
	Get(string, interface{}) (bool, error)
}

// LNclient is an abstraction of a client that connects to a Lightning Network node implementation (like lnd, c-lightning and eclair)
// and provides the methods required by the paywall.
type LNclient interface {
	// GenerateInvoice generates a new invoice based on the price in Satoshis and with the given memo.
	GenerateInvoice(int64, string) (ln.Invoice, error)
	// CheckInvoice checks if the invoice was settled, given an LN node implementation dependent ID.
	// For example lnd uses the payment hash a.k.a. preimage hash as ID, while Lightning Charge
	// uses a randomly generated string as ID.
	CheckInvoice(string) (bool, error)
}

// invoiceMetaData is data that's required to prevent clients from cheating
// (e.g. have multiple requests executed while having paid only once,
// or requesting an invoice for a cheap endpoint and using the payment proof for an expensive one).
// The type itself is not exported, but the fields have to be (for (un-)marshaling).
type invoiceMetaData struct {
	// The unique identifier for the invoice in the LN node.
	// This is NOT the ID that's used for storing the metadata in the storage.
	// Instead, it's the ID used to retrieve info about an invoice from the LN node.
	// The different implementations use different values as ID, for example
	// lnd uses the payment hash a.k.a. preimage hash as ID, while Lightning Charge
	// uses its own randomly generated string as ID.
	//
	// The ID (or rather *key*) used for storing the metadata in the storage
	// is the payment hash of the invoice, because the client sends the preimage
	// (or in the future also its hash) in the final request and we must be able
	// to look up the metadata with that key.
	ImplDepID string
	Method    string
	Path      string
	Used      bool
}

type frameworkAbstraction interface {
	// getPreimageFromHeader returns the content of the "X-Preimage" header.
	getPreimageFromHeader() string
	// respondWithError sends a response with the given message and status code.
	respondWithError(error, string, int)
	// getHTTPrequest returns a pointer to the current http.Request.
	getHTTPrequest() *http.Request
	// respondWithInvoice sends a response with the given headers, status code and invoice string.
	respondWithInvoice(map[string]string, int, []byte)
	// next moves to the next handler, which might be another middleware or the actual request handler.
	// This method is only called when all previous operations were successful (e.g. the invoice was paid properly).
	// An error only needs to be returned if the specific web framework requires middlewares to return one,
	// like Echo does for example.
	next() error
}

func commonHandler(fa fra