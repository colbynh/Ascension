package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/Ascension/main/common/util"
)

func main() {
	url := flag.String("url", "", "enter a url string")
	post := flag.Bool("post", false, "Set this flag if you are doing a post request.")
	flag.Parse()

	if *post {
		b, err := ioutil.ReadFile("data.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		res, err := util.Post(*url, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Post Response: %v", res)
		return
	}
	res, err := util.Get(*url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response: ", res)
}
