package philips

import (
	"fmt"
	"encoding/json"
	"path/filepath"
	"github.com/ascension/internal/util"
)


// GetAllGroups Returns a list of all groups in the system, 
// each group has a name and unique identification number.
func GetAllGroups(hueURL string) ([]LightGroup, error){
	data, err := util.Get(hueURL)
	if err != nil {
		return nil, err
	}
	groupMap := map[string]LightGroup{}
	err = json.Unmarshal(data, &groupMap)
	if err != nil {
		return nil, err
	}

	groups := []LightGroup{}
	for index, group := range groupMap {
		group.Index = index
		groups = append(groups, group)
	}
	return groups, nil

}

// CreateGroup creates a new group containing the lights 
// specified and optional name.
func CreateGroup(hueURL string, group LightGroup) error {
	bytes, err := json.Marshal(group)
	if err != nil {
		return err
	}

	_, err = util.Post(hueURL, bytes)
	if err != nil {
		return err
	}
	return nil
}

// SetGroupState sets the state of a light based on it'd id.
func SetGroupState(hueURL string, group LightGroup) error {
	stateURI := hueURL+"/"+filepath.Join(group.Index,"action")
	bytes, err := json.Marshal(group.Action)
	if err != nil {
		return err
	}
	_, err = util.Put(stateURI, bytes)
	if err != nil {
		return err
	}
	return nil
}

// GetGroup gets a specific group by name and returns
// a LightGroup struct.
func GetGroup(hueURL, name string) (LightGroup, error) {
	groups, err := GetAllGroups(hueURL)
	if err != nil {
		fmt.Println("Getall error: "+err.Error())
	}

	for _, g := range groups {
		if g.Name == name {
			return g, nil
		}
	}
	return LightGroup{}, fmt.Errorf("Light group: \"%v\" not found", name)
}

// ToggleRoom turns on/off all lights of a given group(room).
func ToggleRoom(hueURL, name string) error {
	lg, err := GetGroup(hueURL, name)
	if err != nil {
		return err
	}

	if lg.Action.On == false {
		lg.Action.On = true
	} else {
		lg.Action.On = false
	}

	err = SetGroupState(hueURL, lg)
	if err != nil {
		return err
	}
	return nil
}