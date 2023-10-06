package write

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	bytestream.UnimplementedByteStreamServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (w *Server) Run(ctx context.Context) error {
	srv := grpc.NewServer()
	bytestream.RegisterByteStreamServer(srv, w)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	go func() {
		_ = srv.Serve(listener)
	}()

	// wait until ctx is done
	select {
	case <-ctx.Done():
		break
	}

	srv.GracefulStop()

	return nil
}

func (w *Server) Read(req *bytestream.ReadRequest, readServer bytestream.ByteStream_ReadServer) error {
	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func (w *Server) Write(writeServer bytestream.ByteStream_WriteServer) error {

	committedSize := 0
	bs := make([]byte, 1000)
	for {
		res, err := writeServer.Recv()

		if err != nil {
			fmt.Printf("got err 3 %v \n", err.Error())
			return err
		}

		fmt.Printf("got value: %v \n", res)
		bs = append(bs, res.Data...)

		if res.GetFinishWrite() {
			committedSize = len(bs)
			break
		}
	}

	return writeServer.SendAndClose(&bytestream.WriteResponse{
		CommittedSize: int64(committedSize),
	})

	// return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
