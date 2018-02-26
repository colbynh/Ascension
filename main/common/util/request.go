package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get request for a url.
func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
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
	body, _ := ioutil.ReadAll(res.Body)
	ret := fmt.Sprintf("%s\n", string(body))
	return ret, nil
}

// Put a json payload to a given url.
func Put(url string, data []byte) (string, error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	ret := fmt.Sprintf("%s\n", string(body))
	return ret, nil
}
