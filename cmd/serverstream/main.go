package main

import (
	"context"

	errorc "github.com/elangreza14/grpc/internal/error"
	firestoreclient "github.com/elangreza14/grpc/internal/firestore"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
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
	errorc.CheckErr(err)
	defer conn.Close()

	firestoreClient := firestore.NewFirestoreClient(conn)

	batchGetDocumentServerStream, err := firestoreclient.NewBatchGetDocument(ctx, firestoreClient)
	errorc.CheckErr(err)

	batchGetDocumentServerStream.Run()
}
