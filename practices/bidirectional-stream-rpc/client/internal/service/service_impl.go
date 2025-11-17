package service

import (
	"context"
	"fmt"
	"log"

	serverApi "github.com/PaneroEduardo/golang-grpc-practices/practices/bidirectional-stream-rpc/server/api"
)

type service struct {
	echoClient serverApi.EchoClient
}

func NewClientService(echoClient serverApi.EchoClient) ClientService {
	return service{
		echoClient: echoClient,
	}
}

func (svc service) GetEchoesMessageFromServer(ctx context.Context, message string, times int) error {
	log.Printf("sending the message to the stream to get the echoes")

	stream, err := svc.echoClient.EchoMessage(ctx)
	if err != nil {
		return fmt.Errorf("error creating the streaming instance: %w", err)
	}

	for i := 0; i < times; i++ {
		req := serverApi.MessageRequest{
			Message: message,
		}
		err := stream.Send(&req)
		if err != nil {
			return fmt.Errorf("error trying to send the value: %w", err)
		}
		res, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("error trying to receiving the value: %w", err)
		}

		message = res.GetMessage()
		log.Printf("message received: %s", message)
	}

	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("error trying to get the response: %w", err)
	}

	return nil
}
