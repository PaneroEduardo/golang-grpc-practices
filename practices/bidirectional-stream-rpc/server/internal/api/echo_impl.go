package api

import (
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "github.com/PaneroEduardo/golang-grpc-practices/practices/bidirectional-stream-rpc/server/api"
)

type EchoServer struct {
	pb.UnimplementedEchoServer
}

func NewEchoServer() EchoServer {
	return EchoServer{}
}

func (svc EchoServer) EchoMessage(streamSvr grpc.BidiStreamingServer[pb.MessageRequest, pb.MessageResponse]) error {
	for {
		req, err := streamSvr.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error receiving the stream request: %w", err)
		}
		log.Printf("request message received %s", req.GetMessage())

		res := pb.MessageResponse{
			Message: fmt.Sprintf("echo %s", req.GetMessage()),
		}
		err = streamSvr.Send(&res)
		if err != nil {
			return fmt.Errorf("error sending the message to the steam: %w", err)
		}
	}

	return nil
}
