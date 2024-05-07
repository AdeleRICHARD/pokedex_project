package pokeapi

type LocationsAreasResponse struct {
	NextUrl *string    `json:"next"`
	PrevUrl *string    `json:"previous"`
	Result  []Location `json:"results"`
}

type LocationInfos struct {
	Result []Pokemon `json:"pokemon"`
}

type Location struct {
	Name *string `json:"name"`
	Url  *string `json:"url"`
}

type Pokemon struct {
	Name *string `json:"name"`
}
