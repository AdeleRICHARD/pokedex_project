package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	NextUrl *string
	PrevUrl *string
}

type Location struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

func NewConfig() *Config {
	return &Config{}
}

func (conf *Config) GetNext() []Location {
	if conf.NextUrl == nil {
		conf.NextUrl = new(string)
		*conf.NextUrl = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(*conf.NextUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	type LocationsArea struct {
		NextUrl string     `json:"next"`
		Result  []Location `json:"results"`
	}

	var locations LocationsArea
	json.NewDecoder(res.Body).Decode(&locations)

	conf.NextUrl = &locations.NextUrl
	return locations.Result
}

func Prev() string {
	return "prev"
}
