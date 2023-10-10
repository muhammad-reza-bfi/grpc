package read

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/genproto/googleapis/bytestream"

	file "github.com/elangreza14/grpc/internal/file"
)

type Client struct {
	client bytestream.ByteStreamClient
	Done   chan bool
}

func NewClient(byteStreamClient bytestream.ByteStreamClient) (*Client, error) {
	return &Client{
		client: byteStreamClient,
		Done:   make(chan bool),
	}, nil
}

func (w *Client) Run(ctx context.Context, fileName string) error {
	// initiate Hitting server with read request
	// and sending file name
	clnt, err := w.client.Read(ctx, &bytestream.ReadRequest{
		ResourceName: fileName,
	})
	if err != nil {
		return err
	}

	// reserve bytes slice data
	bs := make([]byte, 1000)

	//  iterate and receive stream from the server
	for {
		res, err := clnt.Recv()

		// check if stream is end
		if err == io.EOF {
			fmt.Printf("streaming is done")
			break
		}

		if err != nil {
			return err
		}

		fmt.Printf("data from server: %v \n", res.GetData())

		// merging data
		bs = append(bs, res.GetData()...)
	}

	// build file to store response
	res := &file.File{
		Name: fmt.Sprintf("%d-%s", time.Now().Nanosecond(), fileName),
		Data: bs,
	}
	return res.Write("output/read")
}
