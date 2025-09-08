package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List pokemon within a locaiton
func (c *Client) ExploreLocation(location []string) (RespShallowExplore, error) {
	url := baseURL + "/location-area"

	if len(location) != 0 {
		for _, val := range location {
			url = url + "/" + val
		}
	}

	if data, ok := c.cache.Get(url); ok {
		exploreResp := RespShallowExplore{}
		err := json.Unmarshal(data, &exploreResp)
		if err != nil {
			return RespShallowExplore{}, err
		}
		return exploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowExplore{}, err
	}
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowExplore{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowExplore{}, err
	}

	c.cache.Add(url, dat)

	exploreResp := RespShallowExplore{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespShallowExplore{}, err
	}

	return exploreResp, nil
}

/*TODO
	handle argument... do we send as slice of string and deal with here or add argument field to config struct and pass as pointer?

	make request/unmarshal data/cache???


*/

