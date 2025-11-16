package service

import "context"

type ClientService interface {
	GetSumFromServer(ctx context.Context, values []int) error
}
