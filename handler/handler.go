package handler

import (
	"context"
	"github.com/PiotrKowalski/square/config"
	pb "github.com/PiotrKowalski/square/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	*pb.UnimplementedServiceServer
	conf config.Config
}

func NewService(config config.Config) *Service {
	return &Service{
		UnimplementedServiceServer: &pb.UnimplementedServiceServer{},
		conf:                       config,
	}
}

func (s *Service) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {

	switch msg := req.GetMessage(); msg {
	case "echo":
		return s.getEcho()

	case "timestamp":
		return s.getTimestamp(), nil

	case "env":
		return s.getEnvs()

	case "version":
		return s.getVersion()

	default:
		return nil, status.Error(codes.InvalidArgument, "Use one of echo|timestamp|env|version as message")
	}
}
