package write

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/bytestream"
)

type Client struct {
	client bytestream.ByteStream_WriteClient
}

func NewClient(ctx context.Context, byteStreamClient bytestream.ByteStreamClient) (*Client, error) {
	clnt, err := byteStreamClient.Write(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: clnt,
	}, nil
}

func (w *Client) Run(fileName string, payloads ...[]byte) error {
	for i := 0; i < len(payloads); i++ {
		req := &bytestream.WriteRequest{
			ResourceName: fileName,
			FinishWrite:  false,
			Data:         payloads[i],
		}

		fmt.Println("sending:", payloads[i])

		// if last set finish write = true
		if i == len(payloads)-1 {
			req.FinishWrite = true
		}

		err := w.client.Send(req)
		if err != nil {
			return err
		}
	}

	res, err := w.client.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Println(res)

	return nil
}
