
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