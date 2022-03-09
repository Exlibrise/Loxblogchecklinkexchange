package wall

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewGinMiddleware returns a Gin middleware in the form of a gin.HandlerFunc.
func NewGinMiddleware(invoiceOptions InvoiceOptions, lnClient LNclient, storageClient StorageClient) gin.HandlerFunc {
	invoiceOptions = assignDefaultValues(invoiceOptions)
	return func(ctx *gin.Context) {
		fa := ginAbstraction{
			ctx: ctx,
		}
		commonHandler(fa, invoiceOptions, lnClient, storageClient)
	}
}

type ginAbstraction struct {
	ctx *gin.Context
}

func (fa ginAbstraction) getPreimageFromHeader() string {
	return fa.ctx.GetHeader("x-preimage")
}

func (fa ginAbstraction) respondWithError(err error, errorMsg string, statusCode int) {
	http.Error(fa.ctx.Writer, errorMsg, statusCode)
	fa.ctx.Abort()
}

func (fa ginAbstraction) getHTTPrequest() *http.Request {
	return fa.ctx.Request
}

func (fa g