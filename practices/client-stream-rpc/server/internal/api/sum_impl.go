package api

import (
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "github.com/PaneroEduardo/golang-grpc-practices/practices/client-stream-rpc/server/api"
)

type SumServer struct {
	pb.UnimplementedSumServer
}

func NewSumServer() SumServer {
	return SumServer{}
}

func (svc SumServer) SumItems(streamSvr grpc.ClientStreamingServer[pb.SumItemRequest, pb.SumResponse]) error {
	sum := 0
	for {
		req, err := streamSvr.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error receiving the stream request: %w", err)
		}
		log.Printf("request value received %d", req.GetValue())
		sum += int(req.GetValue())
	}

	res := pb.SumResponse{
		Total: int32(sum),
	}
	return streamSvr.SendAndClose(&res)
}
