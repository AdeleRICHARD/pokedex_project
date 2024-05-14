package pokeapi

import (
	"net/http"
	"time"

	"github.com/AdeleRICHARD/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2/"

type Config struct {
	NextUrl *string
	PrevUrl *string
	Pokedex map[string]PokemonInfos
}

type Client struct {
	HttpClient http.Client
	Config     *Config
	Cache      *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		HttpClient: http.Client{
			Timeout: 10 * time.Second,
		},
		Config: &Config{
			Pokedex: make(map[string]PokemonInfos),
		},
		Cache: pokecache.NewCache(time.Hour),
	}
}

func NewConfig() *Config {
	return &Config{}
}
