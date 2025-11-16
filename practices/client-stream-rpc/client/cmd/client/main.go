package main

import (
	"context"
	"log"

	"github.com/PaneroEduardo/golang-grpc-practices/practices/client-stream-rpc/server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/PaneroEduardo/golang-grpc-practices/practices/client-stream-rpc/client/internal/service"
)

const (
	serverAddress = "localhost:50051" // Usa el nombre del servicio si estás en Docker Compose
)

var values = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	ctx := context.Background()

	log.Println("creating the grpc client")
	conn, err := grpc.NewClient(serverAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error trying to get the connection with the server: %v", err)
	}
	defer conn.Close()

	log.Println("creating sumclient")
	sumClient := api.NewSumClient(conn)

	log.Println("creating new client service")
	clientService := service.NewClientService(sumClient)

	err = clientService.GetSumFromServer(ctx, values)
	if err != nil {
		log.Fatalf("error trying to get the greetings from server: %v", err)
	}

}
