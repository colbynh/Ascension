package util

import (
	"encoding/json"
)

// ToJSON serializes a json string to an interface.
func ToJSON(data []byte, i *interface{}) {
	json.Unmarshal(data, &i)
}
