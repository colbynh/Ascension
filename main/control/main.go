package main

import (
	"SDK/philips"
	"fmt"
	"time"
)

func main() {
	// Initial integration test
	allLights, err := lights.GetAll()
	if err != nil {
		fmt.Println("Getall error: "+err.Error())
	}
	time.Sleep(time.Second *2)
	for _, l := range allLights {
		time.Sleep(time.Second *2)
		l.State.On = false
		err = lights.SetState(l)
		fmt.Println(err)
	}
}
