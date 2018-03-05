package main

import (
	"strconv"
	"sdk/philips"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	// Initial integration tests
	allLights, err := lights.GetAll()
	if err != nil {
		fmt.Println("Getall error: "+err.Error())
	}

	allgroups, err := lights.GetAllGroups()
	if err != nil {
		fmt.Println("Getall error: "+err.Error())
	}
	fmt.Printf("%+v\n", allgroups)

	RenameAllLights(allLights)
	RandColors(allLights)
}

func RenameAllLights(lts []lights.ColorHue) {
	for i, l := range lts {
		l.Name = "Colby"+strconv.Itoa(i)
		err := lights.Rename(l)
		if err != nil {
			fmt.Println(err)
		}
		
	}
	newLts, err := lights.GetAll()
	if err != nil {
		fmt.Println("Getall error: "+err.Error())
	}
	
	for _, l := range newLts {
		fmt.Println("New Light: ", l.Name)
	}
}

func RandColors(lts []lights.ColorHue) {
	for i := 0; i < 2; i++ {
		for _, l := range lts {
			l.State.Hue = rand.Intn(100000)
			err := lights.SetState(l)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second *1)
		}
	}
}

