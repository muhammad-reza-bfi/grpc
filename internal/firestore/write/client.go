package write

import (
	"context"
	"errors"
	"fmt"
	"io"

	terminal "github.com/elangreza14/grpc/internal/terminal"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var databaseUrl = "projects/elangreza-golang-base/databases/(default)"

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

		// check if stream is end
		if err == io.EOF {
			fmt.Printf("streaming is done")
			break
		}

		if err != nil {
			return err
		}

		if res.WriteResults != nil {
			fmt.Printf("success: %v\n", res.WriteResults)
		}

		if w.streamID == "" || w.streamToken == nil {
			w.streamID = res.GetStreamId()
			w.streamToken = res.GetStreamToken()
		}
	}

	return nil
}

func (w *Write) sendMessage(payload *firestore.WriteRequest) error {
	err := w.writeClient.Send(payload)

	// check the error code
	if s, ok := status.FromError(err); ok {
		if s.Code() != codes.OK {
			return errors.New("error stream when sending the message")
		}
	}

	return nil
}

func (w *Write) send() error {
	err := w.sendMessage(&firestore.WriteRequest{
		Database: databaseUrl,
	})

	if err != nil {
		return err
	}

	fmt.Println("connected into firebase")

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

			// read input from terminal
			valText, ok := w.terminal.ValText()
			if ok {
				val := make(map[string]*firestore.Value)
				val[valText] = &firestore.Value{
					ValueType: &firestore.Value_StringValue{
						StringValue: valText,
					},
				}

				message := &firestore.WriteRequest{
					Database: databaseUrl,
					StreamId: w.streamID,
					Writes: []*firestore.Write{
						{
							Operation: &firestore.Write_Update{
								Update: &firestore.Document{
									Name:   fmt.Sprintf("%s/documents/a/%s", databaseUrl, valText),
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
