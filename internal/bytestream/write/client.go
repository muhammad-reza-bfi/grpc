package write

import (
	"context"
	"fmt"

	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

type Client struct {
	client bytestream.ByteStream_WriteClient
}

func NewClient(ctx context.Context, byteStreamClient bytestream.ByteStreamClient) (*Client, error) {
	clnt, err := byteStreamClient.Write(ctx, grpc.UseCompressor(gzip.Name))
	if err != nil {
		return nil, err
	}

	return &Client{
		client: clnt,
	}, nil
}

func (w *Client) Run(fileName string, chunks ...[]byte) error {
	// send all the chunks one by one
	for i := 0; i < len(chunks); i++ {
		// setup the data with the chunk
		req := &bytestream.WriteRequest{
			ResourceName: fileName,
			WriteOffset:  0,
			FinishWrite:  false,
			Data:         chunks[i],
		}

		// if last chunk, set finishWrite = true
		if i == len(chunks)-1 {
			req.FinishWrite = true
		}

		fmt.Println("sending:", chunks[i])
		// fmt.Print(".")

		// sending chunk of data
		err := w.client.Send(req)
		if err != nil {
			return err
		}
	}

	// sending last message and wait for the response
	res, err := w.client.CloseAndRecv()
	if err != nil {
		return err
	}

	// printing last message from server
	fmt.Print("\n", res)
	return nil
}
