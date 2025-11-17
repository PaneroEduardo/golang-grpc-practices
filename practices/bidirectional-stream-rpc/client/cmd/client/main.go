package main

import (
	"context"
	"log"

	"github.com/PaneroEduardo/golang-grpc-practices/practices/bidirectional-stream-rpc/server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/PaneroEduardo/golang-grpc-practices/practices/bidirectional-stream-rpc/client/internal/service"
)

const (
	serverAddress = "localhost:50051"
	message       = "Hello World"
	times         = 5
)

func main() {
	ctx := context.Background()

	log.Println("creating the grpc client")
	conn, err := grpc.NewClient(serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error trying to get the connection with the server: %v", err)
	}
	defer conn.Close()

	log.Println("creating echo client")
	sumClient := api.NewEchoClient(conn)

	log.Println("creating new client service")
	clientService := service.NewClientService(sumClient)

	err = clientService.GetEchoesMessageFromServer(ctx, message, times)
	if err != nil {
		log.Fatalf("error trying to get the echos message from server: %v", err)
	}

}
