package main

import (
	"encoding/json"
	"net/http"

	"github.com/JackFazackerley/hackernews-microservice/internal/store"

	"github.com/sirupsen/logrus"
)

func main() {
	client := http.DefaultClient

	resp, err := client.Get("http://127.0.0.1:8080/all")
	if err != nil {
		logrus.WithError(err).Fatal("getting all")
	}
	defer resp.Body.Close()

	var items []store.Item

	if err := json.NewDecoder(resp.Body).Decode(&items); err != nil {
		logrus.WithError(err).Error("decoding body")
	}

	logrus.Infof("%+v", items)
}
