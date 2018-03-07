package lights

import (
	"encoding/json"
	"github.com/ascension/internal/util"
)

const (
	baseGroupURL = "http://localhost/api/newdeveloper/groups"
)

// GetAllGroups Returns a list of all groups in the system, 
// each group has a name and unique identification number.
func GetAllGroups() ([]LightGroup, error){
	data, err := util.Get(baseGroupURL)
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
func CreateGroup(group LightGroup) error {
	bytes, err := json.Marshal(group)
	if err != nil {
		return err
	}

	_, err = util.Post(baseGroupURL, bytes)
	if err != nil {
		return err
	}
	return nil
}