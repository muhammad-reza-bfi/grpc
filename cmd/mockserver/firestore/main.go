package main

import (
	"context"

	errorc "github.com/elangreza14/grpc/internal/error"
	firestoreServer "github.com/elangreza14/grpc/internal/firestore/mockserver"
)

func main() {

	bs, err := firestoreServer.NewMockServer()
	errorc.CheckErr(err)

	err = bs.Run(context.Background())
	errorc.CheckErr(err)

}
