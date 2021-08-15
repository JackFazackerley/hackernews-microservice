package main

import (
	"sync"

	"core/internal/client"
	"core/internal/config"
	"core/internal/store"
	"core/internal/worker/consumer"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()

	store, err := store.NewPostgres(cfg)
	if err != nil {
		log.WithError(err).Fatal("creating new postgres client")
	}
	defer store.Close()

	c := client.New(nil)

	topStoriesChan := make(chan int)
	wg := &sync.WaitGroup{}

	w := consumer.New(c, store, topStoriesChan)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go w.Run(wg)
	}

	topStories, err := c.TopStories()
	if err != nil {
		log.WithError(err).Error("getting top stories")
	}

	for _, topStory := range topStories {
		topStoriesChan <- topStory
	}
	close(topStoriesChan)

	wg.Wait()
}
