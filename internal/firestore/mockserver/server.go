package mockserver

import (
	"context"
	"fmt"
	"io"
	"net"

	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockServer struct {
	firestore.UnimplementedFirestoreServer
	feedback chan bool
	done     chan bool
}

func NewMockServer() (*MockServer, error) {
	return &MockServer{
		feedback: make(chan bool, 1000),
		done:     make(chan bool),
	}, nil
}

func (ms *MockServer) Run(ctx context.Context) error {
	srv := grpc.NewServer()
	firestore.RegisterFirestoreServer(srv, ms)

	listener, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		return err
	}

	go func() {
		_ = srv.Serve(listener)
	}()

	fmt.Println("running grpc firestore mock service at 50052")

	// wait until ctx is done
	select {
	case <-ctx.Done():
		break
	}

	srv.GracefulStop()

	return nil
}

func (ms *MockServer) Write(ws firestore.Firestore_WriteServer) error {
	fmt.Println("got ws")
	go ms.send(ws)

	return ms.receive(ws)
}

func (ms *MockServer) send(ws firestore.Firestore_WriteServer) error {
	streamID := "test"
	err := ws.Send(&firestore.WriteResponse{
		StreamId:    streamID,
		StreamToken: []byte(streamID),
	})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ms.feedback:
			err := ws.Send(&firestore.WriteResponse{
				StreamId:    streamID,
				StreamToken: []byte(streamID),
				WriteResults: []*firestore.WriteResult{
					{
						UpdateTime: timestamppb.Now(),
					},
				},
				CommitTime: timestamppb.Now(),
			})
			if err != nil {
				return err
			}
		case <-ms.done:
			return nil
		}
	}
}

func (ms *MockServer) receive(ws firestore.Firestore_WriteServer) error {
	for {
		res, err := ws.Recv()

		// check if stream is end
		if err == io.EOF {
			fmt.Printf("streaming is done")
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("data from client:", res.GetWrites())

		if len(res.GetWrites()) > 0 {
			ms.feedback <- true
		}
	}

	ms.done <- true

	return nil
}
