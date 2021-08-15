package client

import (
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

type Item struct {
	ID      int       `json:"id,omitempty"`
	Deleted bool      `json:"deleted,omitempty"`
	Type    string    `json:"type,omitempty"`
	By      string    `json:"by,omitempty"`
	Time    time.Time `json:"time,omitempty"`
	Dead    bool      `json:"dead,omitempty"`
	URL     string    `json:"url,omitempty"`
	Score   int       `json:"score,omitempty"`
	Title   string    `json:"title,omitempty"`
}

func (i *Item) UnmarshalJSON(data []byte) error {
	type Alias Item
	aux := &struct {
		Time int64 `json:"time"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := ffjson.Unmarshal(data, &aux); err != nil {
		return err
	}

	i.Time = time.Unix(aux.Time, 0).UTC()
	return nil
}

func (i Item) IsValid() bool {
	return !i.Dead && !i.Deleted
}