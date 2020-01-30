package parser

import (
	"encoding/json"
)

// APIParser is weird
type APIParser interface {
	GetJSONResponse() string
}

// GetJSONResponse does...
func (o OpenWeatherResponse) GetJSONResponse() string {
	out, err := json.Marshal(o)

	if err != nil {
		panic(err)
	}

	return string(out)
}
