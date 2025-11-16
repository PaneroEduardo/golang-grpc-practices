package service

import (
	"context"
	"fmt"
	"log"

	serverApi "github.com/PaneroEduardo/golang-grpc-practices/practices/client-stream-rpc/server/api"
)

type service struct {
	sumClient serverApi.SumClient
}

func NewClientService(sumClient serverApi.SumClient) ClientService {
	return service{
		sumClient: sumClient,
	}
}

func (svc service) GetSumFromServer(ctx context.Context, values []int) error {
	log.Printf("sending the numbers as streaming")

	stream, err := svc.sumClient.SumItems(ctx)
	if err != nil {
		return fmt.Errorf("error creating the streaming instance: %w", err)
	}

	for _, value := range values {
		req := serverApi.SumItemRequest{
			Value: int32(value),
		}
		err := stream.Send(&req)
		if err != nil {
			return fmt.Errorf("error trying to send the value: %w", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("error trying to get the response: %w", err)
	}

	log.Printf("the result of the sum is %d", res.GetTotal())

	return nil
}
