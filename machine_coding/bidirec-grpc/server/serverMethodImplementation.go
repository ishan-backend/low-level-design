package main

import (
	servicepb "../proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

type ServiceServer struct {
	servicepb.UnimplementedStreamingGrpcServer
}

func (s *ServiceServer) BidirectionalStream(stream servicepb.StreamingGrpc_BidirectionalStreamServer) error {
	var totalMessages uint32
	var concatString string
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		concatString += value.GetValue()
		totalMessages++
		fmt.Println("Received value: ", value.GetValue())

		if totalMessages%5 == 0 {
			fmt.Println("Total messages received: ", totalMessages, "Sending string: ", concatString)
			if err := stream.Send(&servicepb.BidirectionalMessageResponse{
				Uuid:  value.GetUuid(),
				Value: concatString,
			}); err != nil {
				return err
			}

			totalMessages = 0
			concatString = ""
		}
	}
}

// ClientSideStream continuously consumes messages from request
func (s *ServiceServer) ClientSideStream(stream servicepb.StreamingGrpc_ClientSideStreamServer) error {
	var totalMessages uint32
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&servicepb.ClientMessageResponse{
				Success: true, // optionally you can send amount of messages consumed internally, once the stream has closed from client side
			})
		}
		if err != nil {
			return err
		}

		fmt.Println(value.GetUuid(), value.GetValue())
		totalMessages++
	}
}

// ServerSideStream continuously produces messages to the request and pushes it to stream
func (s *ServiceServer) ServerSideStream(req *servicepb.ServerSideRequest, stream servicepb.StreamingGrpc_ServerSideStreamServer) error {
	for {
		select {
		case <-stream.Context().Done():
			fmt.Println("cancel context received!!!")
			return status.Error(codes.Canceled, "Stream has ended")
		default:
			time.Sleep(1 * time.Second)

			if err := stream.SendMsg(&servicepb.ServerSideResponse{
				Value: time.Now().String(),
			}); err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}
		}
	}
}
