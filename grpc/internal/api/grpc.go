package api

import (
	"context"

	pb "grpc/internal/proto"
	"grpc/internal/store"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	pb.UnimplementedAPIServer
	reader store.Reader
}

func New(reader store.Reader) *Server {
	return &Server{
		reader: reader,
	}
}

func (s Server) All(context.Context, *emptypb.Empty) (*pb.Response, error) {
	items, err := s.reader.All()
	if err != nil {
		return &pb.Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Response{Items: items}, nil
}

func (s Server) Jobs(context.Context, *emptypb.Empty) (*pb.Response, error) {
	items, err := s.reader.Jobs()
	if err != nil {
		return &pb.Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Response{Items: items}, nil
}

func (s Server) Stories(context.Context, *emptypb.Empty) (*pb.Response, error) {
	items, err := s.reader.Stories()
	if err != nil {
		return &pb.Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.Response{Items: items}, nil
}
