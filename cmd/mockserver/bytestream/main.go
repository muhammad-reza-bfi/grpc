package main

import (
	"context"

	bytestreamServer "github.com/elangreza14/grpc/internal/bytestream/mockserver"
	errorc "github.com/elangreza14/grpc/internal/error"
)

func main() {
	bs, err := bytestreamServer.NewMockServer()
	errorc.CheckErr(err)

	err = bs.Run(context.Background())
	errorc.CheckErr(err)
}
