package write

import (
	"context"
	"fmt"

	errorc "github.com/elangreza14/grpc/internal/error"
	"google.golang.org/genproto/googleapis/bytestream"
)

type Client struct {
	client bytestream.ByteStream_WriteClient
}

func NewClient(ctx context.Context, byteStreamClient bytestream.ByteStreamClient) (*Client, error) {
	clnt, err := byteStreamClient.Write(ctx)
	errorc.CheckErr(err)

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

		// if last set finish write = true
		if i == len(payloads)-1 {
			req.FinishWrite = true
		}

		err := w.client.Send(req)
		errorc.CheckErr(err)
	}

	res, err := w.client.CloseAndRecv()
	errorc.CheckErr(err)

	fmt.Println(res)

	return nil
}
