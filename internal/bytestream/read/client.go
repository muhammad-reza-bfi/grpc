package read

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"

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
	}, grpc.UseCompressor(gzip.Name))
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
		} else if err != nil {
			return err
		}

		fmt.Printf("data from server: %v \n", res.GetData())
		// fmt.Printf(".")

		// merging data
		bs = append(bs, res.GetData()...)
	}

	// build file to store response
	return w.writeIntoFile(bs, fileName)
}

func (w *Client) writeIntoFile(bs []byte, fileName string) error {
	// build file to store response
	res := &file.File{
		Name: fmt.Sprintf("%d-%s", time.Now().Nanosecond(), fileName),
		Data: bs,
	}
	return res.Write("output/read")
}
