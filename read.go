package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"time"
)

// client is an http client for invoking read requests for the image
var client = &http.Client{
	Timeout: 3 * time.Second,
}

// read reads the image from a given URL
func read(url string) (image.Image, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("err problem creating http request to read the image: %s", err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("err problem executing the http request to read the image: %s", err.Error())
	}

	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err problem reading the image from the content: %s", err.Error())
	}

	return m, nil
}
