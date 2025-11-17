package service

import "context"

type ClientService interface {
	GetEchoesMessageFromServer(ctx context.Context, message string, times int) error
}
