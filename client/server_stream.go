package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hmuir28/go-grpc-chat/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started...")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)

	if err != nil {
		log.Fatalf("could not send names : %v", names)
	}

	for {
		message, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while streaming %v", err)
		}

		log.Println(message)
	}

	log.Printf("streaming finished...")
}
