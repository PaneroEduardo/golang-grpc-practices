package api

import (
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "github.com/PaneroEduardo/golang-grpc-practices/practices/server-stream-rpc/server/api"
)

type HiServer struct {
	pb.UnimplementedHiServer
}

func NewHelloServer() HiServer {
	return HiServer{}
}

func (svc HiServer) HelloWorld(req *pb.HelloRequest, srv grpc.ServerStreamingServer[pb.HelloResponse]) error {
	log.Printf("received hello request for name %s", req.GetName())

	log.Printf("send by gRPC the greeting %d times", req.GetTimes())

	for i := 0; i < int(req.GetTimes()); i++ {
		response := pb.HelloResponse{
			Message: fmt.Sprintf("Hello %s for %d time", req.GetName(), i+1),
		}

		if err := srv.Send(&response); err != nil {
			log.Fatalf("error trying to send the response: %v", err)
			return fmt.Errorf("error trying to send the response: %w", err)
		}
	}

	return nil
}
