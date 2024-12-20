package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (RespExploreLocation, error) {
	url := baseURL + "/location-area/" + location

	if cachedData, exists := c.cache.Get(url); exists {
		var cachedResp RespExploreLocation
		if err := json.Unmarshal(cachedData, &cachedResp); err != nil {
			return RespExploreLocation{}, err
		}
		return cachedResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocation{}, err
	}

	exploreLocatioResp := RespExploreLocation{}
	if err = json.Unmarshal(data, &exploreLocatioResp); err != nil {
		return RespExploreLocation{}, err
	}

	c.cache.Add(url, data)
	return exploreLocatioResp, nil
}
