
package storage

import (
	"github.com/go-redis/redis"
)

// RedisClient is a StorageClient implementation for Redis.
type RedisClient struct {
	c *redis.Client
}

// Set stores the given object for the given key.
func (c RedisClient) Set(k string, v interface{}) error {
	// First turn the passed object into something that Redis can handle
	// (the Set method takes an interface{}, but the Get method only returns a string,