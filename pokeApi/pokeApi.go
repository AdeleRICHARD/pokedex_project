package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AdeleRICHARD/pokedexcli/internal"
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

	type LocationsArea struct {
		NextUrl string     `json:"next"`
		PrevUrl string     `json:"previous"`
		Result  []Location `json:"results"`
	}

	var locations LocationsArea
	cache := internal.NewCache(5 * time.Second)
	values, ok := cache.Get(*conf.NextUrl)
	if ok {
		json.Unmarshal(values, &locations)
		return locations.Result
	}

	res, err := http.Get(*conf.NextUrl)
	if err != nil {
		fmt.Println(err)
	}

	rawData, err := json.Marshal(locations)
	if err != nil {
		fmt.Println(err)
	}

	cache.Add(*conf.NextUrl, rawData)

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(&locations)

	conf.NextUrl = &locations.NextUrl
	if locations.PrevUrl != "" {
		conf.PrevUrl = &locations.PrevUrl
	}

	return locations.Result
}

func (conf *Config) GetPrev() []Location {
	println(conf.PrevUrl)
	if conf.PrevUrl == nil {
		return nil
	}

	res, err := http.Get(*conf.PrevUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	type LocationsArea struct {
		NextUrl string     `json:"next"`
		PrevUrl string     `json:"previous"`
		Result  []Location `json:"results"`
	}

	var locations LocationsArea
	json.NewDecoder(res.Body).Decode(&locations)
	if locations.PrevUrl != "" {
		conf.PrevUrl = &locations.PrevUrl
	}

	return locations.Result
}
