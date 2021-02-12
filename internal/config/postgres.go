package config

type Postgres interface {
	PostgresUsername() string
	PostgresPassword() string
	PostgresDatabase() string
	PostgresAddress() string
}

func (c Config) PostgresUsername() string {
	return c.GetString("postgres.username")
}

func (c Config) PostgresPassword() string {
	return c.GetString("postgres.password")
}

func (c Config) PostgresDatabase() string {
	return c.GetString("postgres.database")
}

func (c Config) PostgresAddress() string {
	return c.GetString("postgres.address")
}
