package models

import (
	"encoding/json"
	"math"
)

type Temperature int

func (t *Temperature) UnmarshalJSON(data []byte) error {
	var tempFloat float64
	if err := json.Unmarshal(data, &tempFloat); err != nil {
		return err
	}
	*t = Temperature(math.Round(tempFloat))
	return nil
}

type WeatherData struct {
	Name    string  `json:"name"`
	Main    Main    `json:"main"`
	Weather Wheater `json:"weather"`
}

type Main struct {
	Celsius  Temperature `json:"temp"`
	Humidity int         `json:"humidity"`
}

type Wheater []struct {
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
