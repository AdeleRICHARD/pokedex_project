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

func (c *Client) GetLocationInfo(name *string) (LocationInfos, error) {
	var locationInfos LocationInfos
	if dataCache, ok := c.Cache.Get(*name); ok {
		err := json.Unmarshal(dataCache, &locationInfos)
		if err != nil {
			return LocationInfos{}, err
		}
		return locationInfos, nil
	}

	endpoint := "location-area/" + *name
	fullUrl := baseUrl + endpoint

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return LocationInfos{}, err
	}

	res, err := c.HttpClient.Do(req)
	if err != nil {
		return LocationInfos{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationInfos{}, fmt.Errorf("bad status code %d, url: %s", res.StatusCode, fullUrl)
	}

	dataInfo, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationInfos{}, err
	}

	c.Cache.Add(fullUrl, dataInfo)

	err = json.Unmarshal(dataInfo, &locationInfos)
	if err != nil {
		return LocationInfos{}, err
	}

	return locationInfos, nil
}
