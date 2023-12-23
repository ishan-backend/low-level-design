package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"

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
	bidirStreamClient, err := rpcClient.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}

	go func() {
		for i := 0; i < 50; i++ {
			err = bidirStreamClient.Send(&servicepb.BidirectionalMessageRequest{
				Uuid:  string(rand.Int31()),
				Value: string('a' + rand.Int31()%21),
			})
		}

		if err = bidirStreamClient.CloseSend(); err != nil { // closing the streaming when non-nil error is received from server
			log.Fatalf("CloseSend acknowledged: %v", err)
		}
	}()

	for {
		res, err := bidirStreamClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Closing connection: %v", err)
		}

		fmt.Println("Response to client: ", res.Value)
	}
}
