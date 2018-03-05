package lights

import (
	"encoding/json"
	"github.com/ascension/internal/util"
)

const (
	baseGroupURL = "http://localhost/api/newdeveloper/groups"
)

// GetAllGroups gets all the hue groups on the network
// and returns any errors.
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