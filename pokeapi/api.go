package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchPokeAPI(url string, count int) (*LocationAPIResponse, int, error) {
	resp, err := http.Get(
		url,
	)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	var jsonResp LocationAPIResponse
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return nil, 0, err
	}
	count = count + 1
	return &jsonResp, count, nil

}
