package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// SetStackEnv sets an enviroment variable 'stack' to
// to a specified stack.
func SetStackEnv(s string) error {
	stacks := map[string]bool{
		"dev": true,
		"stag":   true,
		"prod": true,
	}
	if stacks[s] {
		err := os.Setenv("stack", s)
		if err != nil {
			return fmt.Errorf("util.SetStack error: stack \"%s\" not allowed", s)
		}
		fmt.Println("Running enviroment as: ", os.Getenv("stack"))
		return nil
	}
	return fmt.Errorf("util.SetStack error: stack \"%s\", not found", s) 
}

// GetNetworkIPs returns a slice of all the ip addresses on the
// current network.
func GetNetworkIPs() ([]string, error) {
	var ips []string
	out, err := exec.Command("/bin/sh", "./hack/network.sh").Output()
    if err != nil {
        return ips, err
	}
	ips = strings.Split(strings.TrimSpace(string(out)), "\n")
	return ips, nil
}