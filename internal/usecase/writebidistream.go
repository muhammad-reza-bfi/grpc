package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"

	firestore "github.com/elangreza14/grpc/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WriteBidiStream struct {
	terminal    Terminal
	streamId    string
	streamToken []byte
	writeClient firestore.Firestore_WriteClient
}

func NewWriteBidiStream(ctx context.Context, firestoreClient firestore.FirestoreClient) (*WriteBidiStream, error) {
	writeClient, err := firestoreClient.Write(ctx)
	if err != nil {
		return nil, err
	}
	return &WriteBidiStream{
		terminal:    *NewTerminal(),
		streamId:    "",
		streamToken: nil,
		writeClient: writeClient,
	}, nil
}

func (w *WriteBidiStream) Run() error {

	go w.receive()

	return w.send()
}

func (w *WriteBidiStream) receive() error {
	for {
		res, err := w.writeClient.Recv()

		if sts, ok := status.FromError(err); ok && sts.Code() == codes.Canceled {
			fmt.Printf("got err 1 %v \n", sts.Code())
			return fmt.Errorf("got error %v", sts.Code())
		} else if err == io.EOF {
			fmt.Printf("got err 2 %v \n", err.Error())
			return errors.New("stream closed")
		} else if err != nil {
			fmt.Printf("got err 3 %v \n", err.Error())
			return err
		}

		fmt.Printf("got commit time: %v \n", res.CommitTime)

		if w.streamId == "" || w.streamToken == nil {
			w.streamId = res.GetStreamId()
			w.streamToken = res.GetStreamToken()
		}
	}
}

func (w *WriteBidiStream) sendMessage(payload *firestore.WriteRequest) error {
	err := w.writeClient.Send(payload)

	// check the error code
	if s, ok := status.FromError(err); ok {
		if s.Code() != codes.OK {
			fmt.Printf("got error: %v", s.Code())
			return errors.New("error stream when sending the message")
		}
	}

	return nil
}

func (w *WriteBidiStream) send() error {
	err := w.sendMessage(&firestore.WriteRequest{
		Database: "projects/elangreza-golang-base/databases/(default)",
	})

	if err != nil {
		return err
	}

	for {
		select {
		case <-w.writeClient.Context().Done():
			fmt.Printf("stream ended by server \n")
			return nil
		default:
			// if streamId or streamToken is empty continue
			if w.streamId == "" || w.streamToken == nil {
				continue
			}

			valText, ok := w.terminal.ValText()
			if ok {
				val := make(map[string]*firestore.Value)
				val[valText] = &firestore.Value{
					ValueType: &firestore.Value_StringValue{
						StringValue: valText,
					},
				}

				message := &firestore.WriteRequest{
					Database: "projects/elangreza-golang-base/databases/(default)",
					StreamId: w.streamId,
					Writes: []*firestore.Write{
						{
							Operation: &firestore.Write_Update{
								Update: &firestore.Document{
									Name:   fmt.Sprintf("projects/elangreza-golang-base/databases/(default)/documents/a/%s", valText),
									Fields: val,
								}},
						}},
					StreamToken: w.streamToken,
				}

				err := w.sendMessage(message)

				if err != nil {
					return err
				}
			}
		}
	}
}
