package client

import (
	"fmt"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type Client struct {
	client *http.Client
}

const (
	hnURL      = "https://hacker-news.firebaseio.com/v0"
	topStories = "topstories.json"
	hnItem     = "item"
)

func New(c *http.Client) *Client {
	if c == nil {
		c = http.DefaultClient
	}

	return &Client{client: c}
}

func (c Client) do(path, method string, v interface{}) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", hnURL, path), nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected %d status, got: %d", http.StatusOK, resp.StatusCode)
	}

	if err := ffjson.NewDecoder().DecodeReader(resp.Body, &v); err != nil {
		return fmt.Errorf("decoding response: %w", err)
	}

	return nil
}

func (c Client) TopStories() ([]int, error) {
	var ids []int

	if err := c.do(topStories, http.MethodGet, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func (c Client) Item(id int) (*Item, error) {
	var item *Item

	if err := c.do(fmt.Sprintf("%s/%d.json", hnItem, id), http.MethodGet, &item); err != nil {
		return nil, err
	}

	return item, nil
}
