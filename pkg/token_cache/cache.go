package token_cache

//let's try new stuff
// redis would be better
// map[string]interface{} - simpler
//but i want to test this
import (
	"medtest/config"
	"time"

	"github.com/patrickmn/go-cache"
)

// when to clean cache from expired tokens
const DefaultCleanUp = cache.NoExpiration

type TokenCache struct {
	Cache *cache.Cache
}

func NewCache(Config *config.Config) *TokenCache {
	exp := time.Duration(Config.Cache.Duration) * time.Minute
	if Config.Cache.Duration == -1 {
		exp = cache.NoExpiration
	}
	C := cache.New(exp, DefaultCleanUp)
	return &TokenCache{Cache: C}
}

func (TokenCache *TokenCache) Set(key string, value interface{}) {
	TokenCache.Cache.Set(key, value, cache.DefaultExpiration)
}

func (TokenCache *TokenCache) Get(key string) interface{} {
	value, found := TokenCache.Cache.Get(key)
	if !found {
		return nil
	}
	return value
}
