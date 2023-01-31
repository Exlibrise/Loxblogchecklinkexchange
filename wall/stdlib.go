package wall

import (
	"net/http"
)

// NewHandlerFuncMiddleware returns a function which you can use within an http.HandlerFunc chain.
func NewHandlerFuncMiddleware(invoiceOptions InvoiceOptions, lnClient LNclient, storageClient StorageClient) func(http.HandlerFunc) http.HandlerFunc {
	return func(ne