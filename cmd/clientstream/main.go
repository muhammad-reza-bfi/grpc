package main

import (
	"context"
	"fmt"

	bytestreamWrite "github.com/elangreza14/grpc/internal/bytestream/write"
	errorc "github.com/elangreza14/grpc/internal/error"
	file "github.com/elangreza14/grpc/internal/file"
	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	res, err := file.Read("large.txt")
	errorc.CheckErr(err)

	fmt.Println("uploading", res.Name)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errorc.CheckErr(err)
	defer conn.Close()

	bsc := bytestream.NewByteStreamClient(conn)

	writeClient, err := bytestreamWrite.NewClient(context.Background(), bsc)
	errorc.CheckErr(err)

	err = writeClient.Run(res.Name, res.CreateChunk()...)
	errorc.CheckErr(err)
}
