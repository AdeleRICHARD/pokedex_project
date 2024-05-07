package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(nextUrl *string) (LocationsAreasResponse, error) {
	endpoint := "location-area"
	fullURL := baseUrl + endpoint

	if nextUrl != nil {
		fullURL = *nextUrl
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationsAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationsAreasResponse{}, fmt.Errorf("bad status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsAreasResponse{}, err
	}

	locationAreasRes := LocationsAreasResponse{}
	err = json.Unmarshal(data, &locationAreasRes)
	if err != nil {
		return LocationsAreasResponse{}, err
	}

	return locationAreasRes, nil
}
