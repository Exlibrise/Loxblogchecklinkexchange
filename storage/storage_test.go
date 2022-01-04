
package storage_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/philippgille/ln-paywall/wall"
)

type foo struct {
	Bar string
}

// testStorageClient tests if reading from and writing to the storage works properly.
func testStorageClient(storageClient wall.StorageClient, t *testing.T) {
	key := strconv.FormatInt(rand.Int63(), 10)

	// Initially the key shouldn't exist
	found, err := storageClient.Get(key, new(foo))
	if err != nil {
		t.Error(err)
	}
	if found {
		t.Errorf("A value was found, but no value was expected")
	}

	// Store an object
	val := foo{