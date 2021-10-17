package storage_test

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
)

// TestBoltClientImpl tests if the BoltClient struct implements the StorageClient interface.
// This doesn't happen at runtime, but at compile time.
func TestBoltClientImpl(t *testing.T) {
	t.SkipNow()
	invoiceOptions := wall.Invoi