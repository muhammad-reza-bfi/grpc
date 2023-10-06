package main

import (
	"context"
	"fmt"

	bytestreamWrite "github.com/elangreza14/grpc/internal/bytestream/write"
	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	bs := []byte("test send text")

	var divided [][]byte
	var numCPU = 3

	chunkSize := (len(bs) + numCPU - 1) / numCPU

	for i := 0; i < len(bs); i += chunkSize {
		end := i + chunkSize

		if end > len(bs) {
			end = len(bs)
		}

		divided = append(divided, bs[i:end])
	}

	// fmt.Printf("%#v\n", divided)
	// res := make([]byte, stat.Size())
	// for i := 0; i < len(divided); i++ {
	// 	for j := 0; j < len(divided[i]); j++ {
	// 		res = append(res, divided[i][j])
	// 	}
	// }

	str := string(bs)
	fmt.Println(str)

	// str2 := string(res)
	// fmt.Println(str2)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bsc := bytestream.NewByteStreamClient(conn)

	writeClient, err := bytestreamWrite.NewClient(context.Background(), bsc)
	if err != nil {
		panic(err)
	}

	err = writeClient.Run(divided...)
	if err != nil {
		panic(err)
	}
}
