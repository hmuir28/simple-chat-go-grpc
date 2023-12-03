package main

import (
	"log"
	"net"
	pb "github.com/hmuir28/go-grpc-chat/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb

	err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

}
