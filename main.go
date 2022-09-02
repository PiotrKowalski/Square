package main

import (
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

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	handler := handler.New()

	s := grpc.NewServer()

	pb.RegisterServiceServer(s, handler)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	//err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)

}
