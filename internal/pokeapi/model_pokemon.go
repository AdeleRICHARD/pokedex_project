package pokeapi

type PokemonInfos struct {
	Exp    int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	PokemonStats
}

type PokemonStats struct {
	BaseStat int    `json:"base_stat"`
	Stat     []Stat `json:"stat"`
	Types    []Type `json:"types"`
}

type Stat struct {
	Name *string `json:"name"`
}

type Type struct {
	Slot int `json:"slot"`
	Type struct {
		Name *string `json:"name"`
	} `json:"type"`
}
