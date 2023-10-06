package write

import (
	"context"
	"fmt"
	"net"
	"time"

	errorc "github.com/elangreza14/grpc/internal/error"
	file "github.com/elangreza14/grpc/internal/file"
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

	listener, err := net.Listen("tcp", "localhost:50051")
	errorc.CheckErr(err)

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

	readServer.Send(&bytestream.ReadResponse{
		Data: []byte{},
	})

	return status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func (w *Server) Write(writeServer bytestream.ByteStream_WriteServer) error {

	committedSize := 0
	bs := make([]byte, 1000)
	for {
		res, err := writeServer.Recv()

		errorc.CheckErr(err)

		fmt.Printf("got value: %v for %v \n", res.Data, res.ResourceName)
		bs = append(bs, res.Data...)
		committedSize = committedSize + len(res.Data)

		if res.GetFinishWrite() {
			err := file.Write(&file.File{
				Name: fmt.Sprintf("%d-%s", time.Now().Nanosecond(), res.ResourceName),
				Data: bs,
			}, "output")

			errorc.CheckErr(err)

			fmt.Println("finished writing:", res.GetResourceName())
			break
		}
	}

	return writeServer.SendAndClose(&bytestream.WriteResponse{
		CommittedSize: int64(committedSize),
	})
}
