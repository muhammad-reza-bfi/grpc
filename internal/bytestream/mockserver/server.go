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
	_ "google.golang.org/grpc/encoding/gzip" // Install the gzip compressor
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

func (w *MockServer) Read(
	req *bytestream.ReadRequest,
	readMockServer bytestream.ByteStream_ReadServer,
) error {
	fmt.Println("querying", req.GetResourceName())

	// read file and search by resource name
	res, err := file.Read(req.GetResourceName())
	if err != nil {
		return status.Errorf(codes.NotFound, "not found")
	}

	// create chunk of bytes from entire file
	for _, data := range res.CreateChunk() {
		fmt.Println("sending:", data)
		// fmt.Print(".")

		err := readMockServer.Send(&bytestream.ReadResponse{
			Data: data,
		})
		if err != nil {
			return status.Errorf(codes.DataLoss, "data loss")
		}
	}

	fmt.Println("data's sent")
	return nil
}

func (w *MockServer) Write(writeMockServer bytestream.ByteStream_WriteServer) error {
	// prepare var
	fileName := ""
	bytes := make([]byte, 1000) // for saving the chunks
	fmt.Printf("\n\ngot write stream\n")
	for {

		// receive payload from client
		res, err := writeMockServer.Recv()

		// check if stream is end
		if err == io.EOF {
			fmt.Printf("streaming is done")
			break
		} else if err != nil {
			return err
		}

		fmt.Println("data from client:", res.Data)
		// fmt.Print(".")

		// appending chunks into complete data
		bytes = append(bytes, res.Data...)

		// check if client want to end the stream, and set the filename
		if res.GetFinishWrite() {
			fileName = res.ResourceName
			break
		}
	}

	// write bytes into file
	err := w.writeIntoFile(bytes, fileName)
	if err != nil {
		return err
	}
	fmt.Print("\nfinished writing:", fileName)

	// send last message and send the committedSize, and the close the stream
	return writeMockServer.SendAndClose(&bytestream.WriteResponse{
		CommittedSize: int64(len(bytes)),
	})
}

func (w *MockServer) writeIntoFile(data []byte, res string) error {
	// write into new file
	fw := &file.File{
		Name: fmt.Sprintf("%d-%s", time.Now().Nanosecond(), res),
		Data: data,
	}
	// place under Output directory
	err := fw.Write("output/write")
	if err != nil {
		return err
	}

	return nil
}
