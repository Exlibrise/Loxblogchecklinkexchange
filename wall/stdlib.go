package wall

import (
	"net/http"
)

// NewHandlerFuncMiddleware returns a function which you can use within an http.HandlerFunc chain.
func New