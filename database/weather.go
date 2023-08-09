package database

import (
	"MT-GO/structs"
	"MT-GO/tools"
	"encoding/json"
)

var weather = structs.Weather{}

func GetWeather() *structs.Weather {
	return &weather
}
func setWeather() {
	raw := tools.GetJSONRawMessage(weatherPath)
	err := json.Unmarshal(raw, &weather)
	if err != nil {
		panic(err)
	}
}
