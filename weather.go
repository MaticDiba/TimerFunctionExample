package main

import "log"

type WeatherDataResponse []WeatherData

type WeatherData struct {
	Mode   string `json:"mode"`
	Path   string `json:"path"`
	Date   string `json:"date"`
	HHMM   string `json:"hhmm"`
	BBox   string `json:"bbox"`
	Width  string `json:"width"`
	Height string `json:"height"`
	Valid  string `json:"valid"`
}

func (wd WeatherDataResponse) Print() {
	for _, data := range wd {
		log.Printf("%+v`\n", data)
	}
}
