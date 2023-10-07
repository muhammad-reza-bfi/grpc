package main

import (
	"context"

	bytestreamRead "github.com/elangreza14/grpc/internal/bytestream/read"
	errorc "github.com/elangreza14/grpc/internal/error"
	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errorc.CheckErr(err)
	defer conn.Close()

	bsc := bytestream.NewByteStreamClient(conn)
	readClient, err := bytestreamRead.NewClient(bsc)
	errorc.CheckErr(err)

	err = readClient.Run(context.Background(), "test.txt")
	errorc.CheckErr(err)
}
