package main

import (
	"fmt"
	"net"

	"grpc/internal/api"
	"grpc/internal/config"
	pb "grpc/internal/proto"
	"grpc/internal/store"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	port = 50051
)

func main() {
	cfg := config.New()

	store, err := store.NewPostgres(cfg)
	if err != nil {
		log.WithError(err).Fatal("creating new postgres client")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.WithError(err).Fatal("starting listen")
	}

	server := api.New(store)

	gs := grpc.NewServer()
	pb.RegisterAPIServer(gs, server)
	if err := gs.Serve(lis); err != nil {
		log.WithError(err).Error("serve")
	}
}
