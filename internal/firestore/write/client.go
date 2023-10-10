package write

import (
	"context"
	"fmt"
	"io"

	terminal "github.com/elangreza14/grpc/internal/terminal"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
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
	fmt.Println("client: create ws")
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
		} else if err != nil {
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

func (w *Write) send() error {
	// initiate first message with empty payload to get streamID
	err := w.writeClient.Send(w.buildMessage(nil))
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
				// build message request with payload
				message := w.buildMessage(&valText)

				// send message to server
				err := w.writeClient.Send(message)
				if err != nil {
					return err
				}
			}
		}
	}
}

func (w *Write) buildMessage(valText *string) *firestore.WriteRequest {
	res := &firestore.WriteRequest{
		Database: databaseUrl,
	}

	if w.streamID != "" {
		res.StreamId = w.streamID
	}
	if w.streamToken != nil {
		res.StreamToken = w.streamToken
	}
	if valText != nil {
		res.Writes = []*firestore.Write{
			{
				Operation: &firestore.Write_Update{
					Update: &firestore.Document{
						Name: fmt.Sprintf("%s/documents/a/%s", databaseUrl, *valText),
					},
				},
			},
		}
	}

	return res
}
