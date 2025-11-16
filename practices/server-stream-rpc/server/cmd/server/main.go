package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	grpcApi "github.com/PaneroEduardo/golang-grpc-practices/practices/server-stream-rpc/server/api"
	internalApi "github.com/PaneroEduardo/golang-grpc-practices/practices/server-stream-rpc/server/internal/api"
)

const port = ":50051"

func main() {
	log.Println("starting grpc server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening the port %s: %v", port, err)
	}
	log.Printf("listening grpc server on port %s", port)

	grpcSvr := grpc.NewServer()
	helloServer := internalApi.NewHelloServer()

	grpcApi.RegisterHiServer(grpcSvr, helloServer)

	err = grpcSvr.Serve(lis)
	if err != nil {
		log.Fatalf("error serving the grpc server %v", err)
	}

}
