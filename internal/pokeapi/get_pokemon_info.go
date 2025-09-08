package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//obtain pokemon data and perform catch logic
func (c *Client) GetPokemonInfo(pokemonName []string) (Pokemon, error) {
	url := baseURL + "/pokemon"
	if len(pokemonName) != 0 {
		for _, val := range pokemonName {
			url = url + "/" + val
		}
	}

	if data, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	
	c.cache.Add(url, dat)

	return pokemonResp, nil
}

