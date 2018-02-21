package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get request for a url.
func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	resStr := fmt.Sprintf("%s\n", string(bytes))
	return resStr, nil
}

// Post a json payload to a given url.
func Post(url string, data []byte) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	ret := fmt.Sprintf("%s\n", string(body))
	return ret, nil
}
