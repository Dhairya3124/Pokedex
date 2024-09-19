package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchPokeAPI(url string) (*LocationAPIResponse, error) {
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
