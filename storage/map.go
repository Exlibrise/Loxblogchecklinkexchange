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
	m.m.Store