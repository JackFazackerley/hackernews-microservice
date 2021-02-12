package config

type Cache interface {
	CacheType() string
	Redis
}

func (c Config) CacheType() string {
	return c.GetString("cache.type")
}
