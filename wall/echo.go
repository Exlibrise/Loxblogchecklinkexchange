
package wall

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewEchoMiddleware returns an Echo middleware in the form of an echo.MiddlewareFunc.
func NewEchoMiddleware(invoiceOptions InvoiceOptions, lnClient LNclient, storageClient StorageClient, skipper middleware.Skipper) echo.MiddlewareFunc {
	invoiceOptions = assignDefaultValues(invoiceOptions)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		if skipper == nil {
			skipper = middleware.DefaultSkipper
		}