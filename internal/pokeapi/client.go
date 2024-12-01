package pokeapi

import (
	"net/http"
	"time"
	"github.com/peeta98/pokedex-cli/internal/pokecache"
)

// Client -
type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

// New Client -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
