package ln

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
)

// stdOutLogger logs to stdout, while the default log package loggers log to stderr.
var stdOutLogger = log.New(os.Stdout, "", log.LstdFlags)

// Invoice is a Lightning Network invoice and contains the typical invoice string an