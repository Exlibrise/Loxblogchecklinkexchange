package storage_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
)

// TestGoMapImpl tests if the GoMap struct implements the StorageClient interface.
// This doesn't happen at runtime, but at compile time.
func TestGoMapImpl(t *testing.T) {
	t.SkipNow()
	invoiceOptions := wall.InvoiceOptions{