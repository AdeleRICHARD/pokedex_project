package pokeapi

type LocationsAreasResponse struct {
	NextUrl *string    `json:"next"`
	PrevUrl *string    `json:"previous"`
	Result  []Location `json:"results"`
}

type LocationInfos struct {
	Result []Pokemons `json:"pokemon_encounters"`
}

type Location struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

type Pokemons struct {
	Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name *string `json:"name"`
}
