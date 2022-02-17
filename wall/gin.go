package wall

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewGinMiddleware returns a Gin middleware in the form of a gin.HandlerFunc.
func NewGinMiddleware(invoiceOptions InvoiceOptions, lnClient LNclient, storageClient StorageClient) gin.HandlerFunc {
	invoi