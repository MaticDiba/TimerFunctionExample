package main

import (
	"encoding/json"
	"net/http"
)

func fetchData(url string) (WeatherDataResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data WeatherDataResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
