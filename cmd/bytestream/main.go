package main

import (
	"context"

	bytestreamWrite "github.com/elangreza14/grpc/internal/bytestream/write"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial("firestore.googleapis.com:443", cred)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bytestreamServer, err := bytestreamWrite.NewServer()
	if err != nil {
		panic(err)
	}

	err = bytestreamServer.Run(context.Background())
	if err != nil {
		panic(err)
	}

}
