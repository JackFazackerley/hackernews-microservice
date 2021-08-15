package store

import (
	"io"

	pb "grpc/internal/proto"
)

type DB interface {
	Writer
	Reader

	io.Closer
}

type Writer interface {
	Put(item *pb.Item) error
}

type Reader interface {
	All() ([]*pb.Item, error)
	Jobs() ([]*pb.Item, error)
	Stories() ([]*pb.Item, error)
}
