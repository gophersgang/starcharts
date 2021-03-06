package github

import (
	"github.com/caarlos0/starchart/config"
	"github.com/caarlos0/starchart/internal/cache"
)

// GitHub client struct
type GitHub struct {
	token    string
	pageSize int
	cache    *cache.Redis
}

// New github client
func New(config config.Config, cache *cache.Redis) *GitHub {
	return &GitHub{
		token:    config.GitHubToken,
		pageSize: config.GitHubPageSize,
		cache:    cache,
	}
}
