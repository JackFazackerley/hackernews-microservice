package store

import (
	"io"
	"time"

	"core/internal/client"
)

type DB interface {
	Writer
	Reader

	io.Closer
}

type Writer interface {
	Put(item *client.Item) error
}

type Reader interface {
	All() ([]*Item, error)
	Jobs() ([]*Item, error)
	Stories() ([]*Item, error)
}

// Item defines the reply
type Item struct {
	ID    int64     `json:"id,omitempty" db:"id"`
	Type  string    `json:"type,omitempty" db:"type"`
	By    string    `json:"by,omitempty" db:"by"`
	Time  time.Time `json:"time,omitempty" db:"time"`
	URL   string    `json:"url,omitempty" db:"url"`
	Score int       `json:"score,omitempty" db:"score"`
	Title string    `json:"title,omitempty" db:"title"`
}
