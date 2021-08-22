package pay

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// LNclient is the abstraction of 