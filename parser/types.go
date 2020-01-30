package parser

// It appears that marshalling tries its best even if some of the structs
// are missing or wrong
type reply struct {
	Name string
}

type coord struct {
	Lon float64
	Lat float64
}

type weather struct {
	ID          float64
	Description string
	Main        string
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
	Deg   float64
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
