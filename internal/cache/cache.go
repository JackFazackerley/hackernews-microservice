package cache

import (
	"io"

	"github.com/JackFazackerley/hackernews-microservice/internal/client"
)

type Cache interface {
	Get(key string) (*client.Item, error)
	Put(item *client.Item) error

	io.Closer
}
