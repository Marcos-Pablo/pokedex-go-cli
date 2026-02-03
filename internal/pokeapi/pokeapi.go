package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) ListLocations(pageURL *string) (ShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var data []byte
	data, exists := c.cache.Get(url)

	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return ShallowLocations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return ShallowLocations{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return ShallowLocations{}, err
		}
		c.cache.Add(url, data)
	}

	var locationsResp ShallowLocations
	err := json.Unmarshal(data, &locationsResp)

	if err != nil {
		return ShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) ExploreLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	var data []byte
	data, exist := c.cache.Get(url)

	if !exist {
		res, err := c.httpClient.Get(url)
		if err != nil {
			return Location{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Location{}, err
		}

		c.cache.Add(url, data)
	}

	var res Location
	err := json.Unmarshal(data, &res)

	if err != nil {
		return Location{}, err
	}

	return res, nil
}
