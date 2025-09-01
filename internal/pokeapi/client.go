package pokeapi

import (
	"net/http"
	"time"
	"github.com/jkantner2/pokedexcli/internal/pokecache"
)

//Client -
type Client struct{
	httpClient 	http.Client
	cache 		pokecache.Cache
}

//NewClient -
func NewClient(timeout time.Duration, cacheLife time.Duration) Client {
	return Client{
		cache:	pokecache.NewCache(cacheLife),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
