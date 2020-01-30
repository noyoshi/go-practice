package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// TODO we need to mirror the api json response with some kind of data type object
type reply struct {
	Name string
}

type coord struct {
	Lon float64
	Lat float64
}

type weather struct {
	ID          string
	Description string
	Icon        string
}

type main struct {
	Temp      float64
	FeelsLike float64 `json:"feels_like"` // NOTE - you need this hint to parse something with underscores
	TempMin   float64 `json:"temp_min"`   // for the most part though, you can rely on the auto parsing?
	TempMax   float64 `json:"temp_max"`
	Pressure  float64
	Humidity  float64
}

type wind struct {
	Speed float64
	Def   float64
}

type clouds struct {
	All float64
}

type sys struct {
	Type    float64
	ID      float64
	Country string
	Sunrise float64
	Sunset  float64
}

// OpenWeatherResponse ?
type OpenWeatherResponse struct {
	Coord      coord
	Weather    []weather
	Base       string
	Main       main
	Visibility float64
	Wind       wind
	Clouds     clouds
	DT         float64
	Sys        sys
	Timezone   float64
	ID         float64
	Name       string
	COD        float64
}

// https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/

const apiEndpoint = "https://api.openweathermap.org/data/2.5/weather?zip="

func parseJSON(body []byte) OpenWeatherResponse {
	var openWeatherResponse OpenWeatherResponse
	json.Unmarshal(body, &openWeatherResponse)
	return openWeatherResponse
}

// Runme is a Dank function, I can't believe my linter wants comments
func Runme() {
	apiKey := os.Getenv("WEATHER_KEY")

	if apiKey == "" {
		panic("no api key found for WEATHER_KEY")
	}

	fmt.Println(apiKey)
	url := apiEndpoint + "98105,us&APPID=" + apiKey
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error in api call")
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error in reading body")
	}

	fmt.Println(string(body))
	openWeatherResponse := parseJSON(body)

	fmt.Println(openWeatherResponse)
}
