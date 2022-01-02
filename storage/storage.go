package storage

import (
	"encoding/json"
)

func toJSON(v interface{}) ([]byte, error) {
	retur