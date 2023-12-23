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
	serverStreamResponse, err := rpcClient.ServerSideStream(context.Background(), &servicepb.ServerSideRequest{Uuid: string(rand.Int31())})
	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}

	resp := make([]*string, 0)
	go func(out []*string) {
		for {
			resp, err := serverStreamResponse.Recv()
			if err == io.EOF {
				return
			}

			if err != nil {
				log.Fatalf("Receiving error: %v", err)
				return
			}

			val := resp.GetValue()
			fmt.Println("response got: ", val)
			out = append(out, &val)
		}
	}(resp)

	for {
		select {
		case <-serverStreamResponse.Context().Done():
			fmt.Println("All done, possible error", serverStreamResponse.Context().Err())
			break
		}
	}
}
