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

// InvoiceOptions are the op