package main

import (
	"context"

	bytestreamServer "github.com/elangreza14/grpc/internal/bytestream/server"
	errorc "github.com/elangreza14/grpc/internal/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial("firestore.googleapis.com:443", cred)
	errorc.CheckErr(err)

	defer conn.Close()

	bs, err := bytestreamServer.NewServer()
	errorc.CheckErr(err)

	err = bs.Run(context.Background())
	errorc.CheckErr(err)

}
