package main

import (
	"fmt"
	"github.com/PiotrKowalski/square/config"
	"github.com/PiotrKowalski/square/gateway"
	"github.com/PiotrKowalski/square/handler"
	pb "github.com/PiotrKowalski/square/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"net"
	"os"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	conf, err := config.GetConfig()
	if err != nil {
		grpclog.Fatal(err)
	}

	addr := fmt.Sprintf("localhost:%s", conf.GrpcPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %w", err)
	}

	s := grpc.NewServer()

	pb.RegisterServiceServer(s, handler.NewService(conf))

	// Serve gRPC Server
	grpclog.Info("Serving gRPC on https://", addr)
	go func() {
		grpclog.Fatal(s.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	grpclog.Fatalln(err)
}
