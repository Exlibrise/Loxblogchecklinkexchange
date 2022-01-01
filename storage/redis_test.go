
package storage_test

import (
	"log"
	"strconv"
	"sync"
	"testing"

	"github.com/go-redis/redis"

	"github.com/philippgille/ln-paywall/ln"
	"github.com/philippgille/ln-paywall/storage"
	"github.com/philippgille/ln-paywall/wall"
)

// Don't use the default number ("0"),
// which could lead to valuable data being deleted when a developer accidentally runs the test with valuable data in DB 0.
var testDbNumber = 15 // 16 DBs by default (unchanged config), starting with 0

// TestRedisClientImpl tests if the RedisClient struct implements the StorageClient interface.
// This doesn't happen at runtime, but at compile time.
func TestRedisClientImpl(t *testing.T) {
	t.SkipNow()
	invoiceOptions := wall.InvoiceOptions{}
	lnClient := ln.LNDclient{}
	redisClient := storage.RedisClient{}
	wall.NewHandlerFuncMiddleware(invoiceOptions, lnClient, redisClient)
	wall.NewHandlerMiddleware(invoiceOptions, lnClient, redisClient)
	wall.NewGinMiddleware(invoiceOptions, lnClient, redisClient)
}

// TestRedisClient tests if reading and writing to the storage works properly.
//
// Note: This test is only executed if the initial connection to Redis works.
func TestRedisClient(t *testing.T) {
	if !checkRedisConnection(testDbNumber) {
		t.Skip("No connection to Redis could be established. Probably not running in a proper test environment.")
	}

	deleteRedisDb(testDbNumber) // Prep for previous test runs
	redisOptions := storage.RedisOptions{
		DB: testDbNumber,
	}
	redisClient := storage.NewRedisClient(redisOptions)

	testStorageClient(redisClient, t)
}

// TestRedisClientConcurrent launches a bunch of goroutines that concurrently work with the Redis client.
func TestRedisClientConcurrent(t *testing.T) {