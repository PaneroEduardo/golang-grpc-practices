package service

import "context"

type ClientService interface {
	GetGreetingFromServer(ctx context.Context, name string, times int) error
}
