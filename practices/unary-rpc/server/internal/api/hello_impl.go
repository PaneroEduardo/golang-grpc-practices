package api

import (
	"context"
	"fmt"
	"log"

	pb "github.com/PaneroEduardo/golang-grpc-practices/practices/unary-rpc/server/api"
)

type HiServer struct {
	pb.UnimplementedHiServer
}

func NewHelloServer() HiServer {
	return HiServer{}
}

func (svc HiServer) HelloWorld(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("received hello request for name: %s", req.GetName())
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %s", req.GetName()),
	}, nil
}
