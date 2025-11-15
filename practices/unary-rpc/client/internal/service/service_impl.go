package service

import (
	"context"
	"errors"
	"fmt"

	serverApi "github.com/PaneroEduardo/golang-grpc-practices/practices/unary-rpc/server/api"
)

type service struct {
	hiClient serverApi.HiClient
}

func NewClientService(hiClient serverApi.HiClient) ClientService {
	return service{
		hiClient: hiClient,
	}
}

func (svc service) GetGreetingFromServer(ctx context.Context, name string) error {
	req := serverApi.HelloRequest{
		Name: name,
	}

	res, err := svc.hiClient.HelloWorld(ctx, &req)
	if err != nil {
		return fmt.Errorf("error trying to get the greeting: %w", err)
	}

	if res == nil {
		return errors.New("the message is empty")
	}

	fmt.Println(res.GetMessage())

	return nil
}
