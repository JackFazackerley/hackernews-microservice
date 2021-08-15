package consumer

import (
	"sync"

	"grpc/internal/client"
	"grpc/internal/store"

	log "github.com/sirupsen/logrus"
)

type Worker struct {
	client     *client.Client
	store      store.Writer
	topStories <-chan int
}

func New(c *client.Client, store store.Writer, topStories <-chan int) *Worker {
	return &Worker{
		client:     c,
		store:      store,
		topStories: topStories,
	}
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	for topStory := range w.topStories {
		item, err := w.client.Item(topStory)
		if err != nil {
			log.WithError(err).Error("getting story")
			continue
		}

		if !item.Dead && !item.Deleted {
			if err := w.store.Put(item); err != nil {
				log.WithError(err).Error("putting item into store")
			}
		}
	}
}
