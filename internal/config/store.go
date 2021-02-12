package config

type Store interface {
	StoreType() string
	Postgres
}

func (c Config) StoreType() string {
	return c.GetString("store.type")
}
