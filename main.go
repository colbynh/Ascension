package main

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"github.com/ascension/server"
	"github.com/ascension/internal/util"
	"github.com/ascension/config"
	"flag"
)

var (
	stack = flag.String("stack", "dev", "Set Enviroment Stack")
	philipsDir = "philips"
)

func main() {
	flag.Parse()

	err := util.SetStackEnv(*stack)
	if err != nil {
		panic(err)
	}

	config := initConfigs()
	err = server.Start(config)
	if err != nil {
		panic(err)
	}
}

func initConfigs() config.VendorConfig {
	appdir := filepath.Join(os.Getenv("HOME"), "ascension")
	vendorCfg := config.VendorConfig{}

	// Get philips config
	pb, err := ioutil.ReadFile(filepath.Join(appdir, philipsDir, "config.json"))
	if err != nil {
		panic(err)
	}
	pCfg := config.PhilipsConfig{}
	util.ToJSON(pb, pCfg)
	vendorCfg.Philips = pCfg

	return vendorCfg
}

