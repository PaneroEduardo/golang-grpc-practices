package service

import (
	"context"
	"fmt"
	"io"
	"log"

	serverApi "github.com/PaneroEduardo/golang-grpc-practices/practices/server-stream-rpc/server/api"
)

type service struct {
	hiClient serverApi.HiClient
}

func NewClientService(hiClient serverApi.HiClient) ClientService {
	return service{
		hiClient: hiClient,
	}
}

func (svc service) GetGreetingFromServer(ctx context.Context, name string, times int) error {
	log.Printf("sending the name %s and %d times to print the name", name, times)
	req := serverApi.HelloRequest{
		Name:  name,
		Times: int32(times),
	}

	log.Printf("starting the streaming")
	stream, err := svc.hiClient.HelloWorld(ctx, &req)
	if err != nil {
		return fmt.Errorf("error trying to get the greeting: %w", err)
	}
	for {
		message, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error receiving the message %w", err)
		}

		fmt.Println(message.GetMessage())
	}

	return nil
}
