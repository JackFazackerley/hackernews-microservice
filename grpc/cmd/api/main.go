package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"grpc/internal/config"
	"grpc/internal/handler"
	pb "grpc/internal/proto"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.New()

	conn, err := grpc.Dial(cfg.GRPCAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.WithError(err).Fatal("GRPC dial")
	}
	defer conn.Close()

	client := pb.NewAPIClient(conn)

	apiHandler := handler.New(client)

	r := mux.NewRouter()

	r.Handle("/all", handler.Handler{H: apiHandler.All}).Methods(http.MethodGet)
	r.Handle("/jobs", handler.Handler{H: apiHandler.Jobs}).Methods(http.MethodGet)
	r.Handle("/stories", handler.Handler{H: apiHandler.Stories}).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.WithError(err).Error("starting server")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		if err != http.ErrServerClosed {
			log.WithError(err).Error("shutting down server")
		}
	}

	log.Warn("shutting down")
}
