package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Config struct {
	NextUrl *string
	PrevUrl *string
}

type Client struct {
	httpClient http.Client
	Config     *Config
}

func NewClient() *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
		Config: &Config{},
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (conf *Config) GetPrev(name *string) []Location {
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
