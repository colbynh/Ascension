package philips

import (
	"strings"
	"path/filepath"
	"encoding/json"
	"fmt"
	"github.com/ascension/internal/util"
)

// GetAll gets all the lights on the network
func GetAll(hueURL string) ([]ColorHue, error) {

	data, err := util.Get(hueURL)
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
func SetState(hueURL string, light ColorHue) error {
	stateURI := hueURL+"/"+filepath.Join(light.Index,"state")
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
func Rename(hueURL string, light ColorHue) error {
	deviceURI := hueURL+"/"+light.Index
	rawJSON := []byte(`{"name":"`+light.Name+`"}`)
	_, err := util.Put(deviceURI, rawJSON)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// Delete a light device
func Delete(hueURL, id string) error {
	deviceURI := hueURL+"/"+id
	_, err := util.Delete(deviceURI)
	if err != nil {
		return err
	}
	return nil
}

// GetBridgeIP automatically detects the local ip of the philips hue bridge
func GetBridgeIP() (string, error) {
	const bridgeRes = "unauthorized user"
	ips, err := util.GetNetworkIPs()
    if err != nil {
        return "", err
	}
	for _, ip := range ips {
		url := "http://"+ip+"/api/username"
		b, _ := util.Get(url)
		if strings.Contains(string(b), bridgeRes) {
			return ip, nil
		}
	}
	return "", fmt.Errorf("error: unable to detect philips hue bridge")
}