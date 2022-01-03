package storage

import (
	"encoding/json"
)

func toJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

fun