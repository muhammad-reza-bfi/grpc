package main

import (
	"context"

	firestore "github.com/elangreza14/grpc/googleapis/firestore/v1beta1"
	usecase "github.com/elangreza14/grpc/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx := context.Background()

	headers := metadata.Pairs(
		"Authorization", "Bearer "+"ya29.a0AfB_byDRJ8BTNFSPzipuQbiiH6Kgdh6jjvkpSl08fVUeB5m63LnjZdHAHHLANNM_pkxBbvm5jhFqzHteo-YLO4xxAnP1UBVuGaHchCQ8kazO6w0lcrjkENlNePWlLmvnbxfDfiKxp88_xn6UpcYwJ2qwdIGQrsD-eblcaCgYKARkSARESFQGOcNnCw8DgFBDCfDAR9bqUThSuVA0171",
	)
	ctx = metadata.NewOutgoingContext(ctx, headers)

	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial("firestore.googleapis.com:443", cred)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	firestoreClient := firestore.NewFirestoreClient(conn)

	// writeBidiStream, err := usecase.NewWriteBidiStream(ctx, firestoreClient)
	// if err != nil {
	// 	panic(err)
	// }
	// writeBidiStream.Run()

	batchGetDocumentServerStream, err := usecase.NewBatchGetDocumentServerStream(ctx, firestoreClient)
	if err != nil {
		panic(err)
	}
	batchGetDocumentServerStream.Run()
}
