package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/hmuir28/go-grpc-chat/proto"
)

func callSayHelloBidrectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional streaming started...")

	stream, err := client.SayHelloBidrectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not send names : %v", names)
	}

	waitChannel := make(chan struct{})

	go func() {
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

		close(waitChannel)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while streaming %v", err)
		}

		time.Sleep(2 * time.Second)
	}

	stream.CloseSend()

	<-waitChannel

	log.Printf("streaming finished...")
}
