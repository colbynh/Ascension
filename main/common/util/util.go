package util

import (
	"encoding/json"
)

// ToJSON serializes a json string to an interface.
func ToJSON(data []byte, i interface{}) error {
	err := json.Unmarshal(data, &i)
	if err != nil {
		return err
	}
	return nil
}
