package firestoreclient

import (
	"context"
	"errors"
	"fmt"
	"io"

	errorc "github.com/elangreza14/grpc/internal/error"
	terminal "github.com/elangreza14/grpc/internal/terminal"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Write struct {
	terminal    terminal.Terminal
	streamID    string
	streamToken []byte
	writeClient firestore.Firestore_WriteClient
}

func NewWrite(ctx context.Context, firestoreClient firestore.FirestoreClient) (*Write, error) {
	writeClient, err := firestoreClient.Write(ctx)
	if err != nil {
		return nil, err
	}
	return &Write{
		terminal:    *terminal.New(),
		streamID:    "",
		streamToken: nil,
		writeClient: writeClient,
	}, nil
}

func (w *Write) Run() error {

	go w.receive()

	return w.send()
}

func (w *Write) receive() error {
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

		if w.streamID == "" || w.streamToken == nil {
			w.streamID = res.GetStreamId()
			w.streamToken = res.GetStreamToken()
		}
	}
}

func (w *Write) sendMessage(payload *firestore.WriteRequest) error {
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

func (w *Write) send() error {
	err := w.sendMessage(&firestore.WriteRequest{
		Database: "projects/elangreza-golang-base/databases/(default)",
	})
	errorc.CheckErr(err)

	for {
		select {
		case <-w.writeClient.Context().Done():
			fmt.Printf("stream ended by server \n")
			return nil
		default:
			// if streamId or streamToken is empty continue
			if w.streamID == "" || w.streamToken == nil {
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
					StreamId: w.streamID,
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
				errorc.CheckErr(err)
			}
		}
	}
}
