package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationsAreasResponse, error) {
	locationAreasRes := LocationsAreasResponse{}

	endpoint := "location-area"
	fullURL := baseUrl + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	respCache, ok := c.Cache.Get(fullURL)
	if ok {
		err := json.Unmarshal(respCache, &locationAreasRes)
		if err != nil {
			return LocationsAreasResponse{}, err
		}
		return locationAreasRes, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationsAreasResponse{}, err
	}

	res, err := c.HttpClient.Do(req)
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

	c.Cache.Add(fullURL, data)

	err = json.Unmarshal(data, &locationAreasRes)
	if err != nil {
		return LocationsAreasResponse{}, err
	}

	return locationAreasRes, nil
}
