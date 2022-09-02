package handler

import (
	"context"
	"errors"
	pb "github.com/PiotrKowalski/square/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Service struct {
	*pb.UnimplementedServiceServer
}

func New() *Service {
	return &Service{
		UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
	}
}

func (s *Service) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {

	switch msg := req.GetMessage(); msg {
	case "echo":

	case "timestamp":
		now := timestamppb.Now()
		return &pb.PingResponse{ReturnMessage: &pb.PingResponse_Timestamp{Timestamp: now}}, nil

	case "env":

	case "version":

	default:
		return nil, errors.New("")
	}

	return nil, errors.New("")
}
