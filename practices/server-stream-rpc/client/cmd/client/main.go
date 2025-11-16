package main

import (
	"context"
	"log"

	"github.com/PaneroEduardo/golang-grpc-practices/practices/server-stream-rpc/server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/PaneroEduardo/golang-grpc-practices/practices/server-stream-rpc/client/internal/service"
)

const (
	serverAddress = "localhost:50051" // Usa el nombre del servicio si estás en Docker Compose
	nameToSend    = "PaneroEduardo"
	times         = 10
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

	log.Println("creating hiclient")
	hiClient := api.NewHiClient(conn)

	log.Println("creating new client service")
	clientService := service.NewClientService(hiClient)

	err = clientService.GetGreetingFromServer(ctx, nameToSend, times)
	if err != nil {
		log.Fatalf("error trying to get the greetings from server: %v", err)
	}

}
