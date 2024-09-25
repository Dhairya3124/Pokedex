package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2/"

func FetchPokeAPI(url string) (*LocationAPIResponse, error) {
	if url == "" {
		url = baseURL + "location-area/"
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
	var jsonResp LocationAPIResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return nil, err
	}
	return &jsonResp, nil

}
