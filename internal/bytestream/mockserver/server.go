package mockserver

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	file "github.com/elangreza14/grpc/internal/file"
	"google.golang.org/genproto/googleapis/bytestream"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MockServer struct {
	bytestream.UnimplementedByteStreamServer
}

func NewMockServer() (*MockServer, error) {
	return &MockServer{}, nil
}

func (w *MockServer) Run(ctx context.Context) error {
	srv := grpc.NewServer()
	bytestream.RegisterByteStreamServer(srv, w)

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		return err
	}

	go func() {
		_ = srv.Serve(listener)
	}()

	fmt.Println("running grpc bytestream mock service at 50051")

	// wait until ctx is done
	select {
	case <-ctx.Done():
		break
	}

	srv.GracefulStop()

	return nil
}

func (w *MockServer) Read(req *bytestream.ReadRequest, readMockServer bytestream.ByteStream_ReadServer) error {
	fmt.Println("querying", req.GetResourceName())

	res, err := file.Read(req.GetResourceName())
	if err != nil {
		return status.Errorf(codes.NotFound, "not found")
	}

	for _, data := range res.CreateChunk() {
		fmt.Println("sending:", data)
		readMockServer.Send(&bytestream.ReadResponse{
			Data: data,
		})
	}

	fmt.Println("data's sent")

	return nil
}

func (w *MockServer) Write(writeMockServer bytestream.ByteStream_WriteServer) error {

	committedSize := 0
	bs := make([]byte, 1000)
	for {
		res, err := writeMockServer.Recv()

		// check if stream is end
		if err == io.EOF {
			fmt.Printf("streaming is done")
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("data from client:", res.Data)
		bs = append(bs, res.Data...)
		committedSize = committedSize + len(res.Data)

		if res.GetFinishWrite() {

			// write into new file
			fw := &file.File{
				Name: fmt.Sprintf("%d-%s", time.Now().Nanosecond(), res.ResourceName),
				Data: bs,
			}
			// place under Output directory
			err := fw.Write("output/write")
			if err != nil {
				return err
			}

			fmt.Println("finished writing:", res.GetResourceName())
			break
		}
	}

	return writeMockServer.SendAndClose(&bytestream.WriteResponse{
		CommittedSize: int64(committedSize),
	})
}
