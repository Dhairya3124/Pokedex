package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	pokecache "github.com/Dhairya3124/PokeDex/pokeCache"
)

const baseURL = "https://pokeapi.co/api/v2/"

func FetchPokeAPI(url string, cache *pokecache.Cache) (*LocationAPIResponse, error) {
	if url == "" {
		url = baseURL + "location-area/"
	}
	cacheResp, cacheHit := cache.Get(url)
	if cacheHit {
		results := LocationAPIResponse{}
		err := json.Unmarshal(cacheResp, &results)
		if err != nil {
			return nil, err
		}
		return &results, nil
	}

	resp, err := http.Get(
		url,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cache.Add(url, body)
	var jsonResp LocationAPIResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

}
func FetchPokeExploreAPI(areaName string, cache *pokecache.Cache) (*LocationAPIResponseDetailed, error) {
	
		url := baseURL + "location-area/" + areaName
		cacheResp, cacheHit := cache.Get(url)
	if cacheHit {
		results := LocationAPIResponseDetailed{}
		err := json.Unmarshal(cacheResp, &results)
		if err != nil {
			return nil, err
		}
		return &results, nil
	}
		
	resp, err := http.Get(
		url,
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cache.Add(url, body)
	var jsonResp LocationAPIResponseDetailed
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return nil, err
	}

	return &jsonResp, nil

	
}