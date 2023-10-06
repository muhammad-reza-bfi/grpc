package main

import (
	"context"

	bytestreamServer "github.com/elangreza14/grpc/internal/bytestream/server"
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

	bs, err := bytestreamServer.NewServer()
	if err != nil {
		panic(err)
	}

	err = bs.Run(context.Background())
	if err != nil {
		panic(err)
	}

}
