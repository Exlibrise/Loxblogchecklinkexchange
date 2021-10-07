
package storage

import (
	"sync"

	bolt "github.com/coreos/bbolt"
)

var bucketName = "ln-paywall"

// BoltClient is a StorageClient implementation for bbolt (formerly known as Bolt / Bolt DB).
type BoltClient struct {
	db   *bolt.DB
	lock *sync.Mutex
}

// Set stores the given object for the given key.
func (c BoltClient) Set(k string, v interface{}) error {
	// First turn the passed object into something that Bolt can handle
	data, err := toJSON(v)
	if err != nil {
		return err
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	err = c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(k), data)