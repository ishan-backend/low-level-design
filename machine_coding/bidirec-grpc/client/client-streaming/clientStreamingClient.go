package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"time"

	servicepb "../../proto"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	rpcClient := servicepb.NewStreamingGrpcClient(conn)
	clientStreamClient, err := rpcClient.ClientSideStream(context.Background())
	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}

	for i := 0; i < 50; i++ {
		err = clientStreamClient.Send(&servicepb.ClientMessageRequest{
			Uuid:  string(rand.Int31()),
			Value: string('a' + rand.Int31()%21),
		})
		time.Sleep(100 * time.Millisecond)
	}

	res, err := clientStreamClient.CloseAndRecv()
	if err != nil {
		log.Fatalln("Closing client connection", err)
	}

	fmt.Println("Successfully acknowledged from server? : ", res.Success)
}
