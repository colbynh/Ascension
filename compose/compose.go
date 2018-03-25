package compose

import (
	"github.com/spf13/viper"
	"sdk/philips"
	"fmt"
)

const (
	hueENV = "HUE"
)

func Init() {
	initPhilips()
}

func initPhilips() {
	viper.SetConfigName("config")
	viper.AddConfigPath("~/.hue/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	
	// Set hue bridge env "HUEip=x.x.x.x"
	ip, err := philips.GetBridgeIP()
	viper.SetEnvPrefix(hueENV)
	viper.BindEnv("ip", ip)
}