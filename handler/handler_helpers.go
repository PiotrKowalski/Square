package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	pb "github.com/PiotrKowalski/square/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"net/http"
	"os"
	"regexp"
)

func (s *Service) getTimestamp() *pb.PingResponse {
	now := timestamppb.Now()
	return &pb.PingResponse{ReturnMessage: &pb.PingResponse_Timestamp{Timestamp: now}}
}

func (s *Service) getVersion() (*pb.PingResponse, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	readFile, err := os.Open(fmt.Sprintf("%s/proto/service.proto", dir))
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	r, err := regexp.Compile(`version: "(\d+.\d+)"`)
	if err != nil {
		return nil, err
	}

	var version string
	for fileScanner.Scan() {
		if len(r.FindStringSubmatch(fileScanner.Text())) != 0 {
			version = r.FindStringSubmatch(fileScanner.Text())[1]
			break
		}
	}

	err = readFile.Close()
	if err != nil {
		return nil, err
	}

	return &pb.PingResponse{ReturnMessage: &pb.PingResponse_Version{Version: version}}, nil
}

func (s *Service) getEnvs() (*pb.PingResponse, error) {
	confMarshaled, err := json.Marshal(s.conf)
	if err != nil {
		return nil, err
	}
	return &pb.PingResponse{ReturnMessage: &pb.PingResponse_Env{Env: string(confMarshaled)}}, nil
}

func (s *Service) getEcho() (*pb.PingResponse, error) {
	resp, err := http.Get(s.conf.EchoUrl)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &pb.PingResponse{ReturnMessage: &pb.PingResponse_Echo{Echo: string(body)}}, nil
}
