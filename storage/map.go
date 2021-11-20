package storage

import (
	"sync"
)

// GoMap is a StorageClient implementation for a simple Go sync.Map.
type GoMap struct {
	m