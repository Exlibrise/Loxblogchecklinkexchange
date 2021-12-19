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
	invoiceOptions := wall.InvoiceOptions{}
	lnClient := ln.LNDclient{}
	goMap := storage.GoMap{}
	wall.NewHandlerFuncMiddleware(invoiceOptions, lnClient, goMap)
	wall.NewHandlerMiddleware(invoiceOptions, lnClient, goMap)
	wall.NewGinMiddleware(invoiceOptions, lnClient, goMap)
}

// TestGoMap tests if reading and writing to the storage works properly.
func TestGoMap(t *testing.T) {
	goMap := storage.NewGoMap()

	testStorageClient(goMap, t)
}

// TestGoMapConcurrent launches a bunch of goroutines that concurrently work with one GoMap.
// The GoMap is a sync.Map, so the concurrency should be supported by the used package.
func TestGoMapConcurrent(t *testing.T) {
	goMap := storage.NewGoMap()

	goroutineCount := 1000

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(goroutineCount) // Must be call