package storage

import (
	"sync"
)

// GoMap is a StorageClient implementation for a simple Go sync.Map.
type GoMap struct {
	m *sync.Map
}

// Set stores the given object for the given key.
func (m GoMap) Set(k string, v interface{}) error {
	data, err := toJSON(v)
	if err != nil {
		return err
	}
	m.m.Store(k, data)
	return nil
}

// Get retrieves the stored object for the given key and populates the fields of the object that v points to
// with the values of the retrieved object's values.
func (m GoMap) Get(k string, v interface{}) (bool, error) {
	data, found := m.