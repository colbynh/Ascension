package main

import (
	"github.com/ascension/server"
	"github.com/ascension/compose"
	"github.com/ascension/internal/util"
	"flag"
	"fmt"
	"sdk/philips"
)

var (
	stack = flag.String("stack", "dev", "Set Enviroment Stack")
)

func main() {
	flag.Parse()
	compose.Init()

	err := util.SetStackEnv(*stack)
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}
}

