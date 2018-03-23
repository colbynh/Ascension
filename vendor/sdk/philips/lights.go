package philips

import (
	"path/filepath"
	"encoding/json"
	"fmt"
	"github.com/ascension/internal/util"
)

const (
	baseURL = "http://10.0.0.130/api/KbYFpMci3bdEnLQgxbdXxwfgo8Bxsn8oBYo3w4e1/lights"
)

// GetAll gets all the lights on the network
func GetAll() ([]ColorHue, error) {
	data, err := util.Get(baseURL)
	if err != nil {
		return nil, err
	}
	lightMap := map[string]ColorHue{}
	err = json.Unmarshal(data, &lightMap)
	if err != nil {
		return nil, err
	}

	lights := []ColorHue{}
	for index, light := range lightMap {
		light.Index = index
		lights = append(lights, light)
	}
	return lights, nil
}

// SetState sets the state of a light based on it'd id.
func SetState(light ColorHue) error {
	stateURI := baseURL+"/"+filepath.Join(light.Index,"state")
	bytes, err := json.Marshal(light.State)
	if err != nil {
		return err
	}
	_, err = util.Put(stateURI, bytes)
	if err != nil {
		return err
	}
	return nil
}

// Rename a light.
func Rename(light ColorHue) error {
	deviceURI := baseURL+"/"+light.Index
	rawJSON := []byte(`{"name":"`+light.Name+`"}`)
	_, err := util.Put(deviceURI, rawJSON)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// Delete a light device
func Delete(id string) error {
	deviceURI := baseURL+"/"+id
	_, err := util.Delete(deviceURI)
	if err != nil {
		return err
	}
	return nil
}

// ToggleRoom turns on/off all lights of a given group(room).
func ToggleRoom(name string) error {
	lg, err := GetGroup(name)
	if err != nil {
		return err
	}

	if lg.Action.On == false {
		lg.Action.On = true
	} else {
		lg.Action.On = false
	}

	err = SetGroupState(lg)
	if err != nil {
		return err
	}
	return nil
}