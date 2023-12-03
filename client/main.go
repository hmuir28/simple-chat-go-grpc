package main

import (
	"log"

	pb "github.com/hmuir28/go-grpc-chat/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NameList{
	// 	Names: []string{ "Harry", "Michael", "Brian" }
	// }

	callSayHello(client)

}
