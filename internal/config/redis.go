package config

type Redis interface {
	RedisAddress() string
}

func (c Config) RedisAddress() string {
	return c.GetString("redis.address")
}
