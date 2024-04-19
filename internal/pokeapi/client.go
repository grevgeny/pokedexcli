package pokeapi

import (
	"net/http"
	"time"

	"github.com/grevgeny/pokedexcli/internal/pokecache"
)

type Client struct {
	http.Client
	cache pokecache.Cache
}

func NewClient(cacheTTL time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(cacheTTL),
	}
}
