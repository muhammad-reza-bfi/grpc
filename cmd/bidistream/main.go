package main

import (
	"context"

	errorc "github.com/elangreza14/grpc/internal/error"
	firestoreWrite "github.com/elangreza14/grpc/internal/firestore/write"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// gen token from here
// https://developers.google.com/oauthplayground/?code=4/0AfJohXkE9e9rW6U7oYpAi5f4khtyexclBV8wr6QZImZtRq8IoPv_Q2MuClbhXy8SYbAhwg&scope=https://www.googleapis.com/auth/cloud-platform

func main() {
	ctx := context.Background()

	headers := metadata.Pairs(
		"Authorization", "Bearer "+"ya29.a0AfB_byANXRLYg37HDLD6pd4IRG-mb-zMSPwIx0XgJX04EzYYOebeZ_qPuwlwVCYnEa5S01z5YKTx9MYy-xZLFUakVLxv6AoSQ6JN5XPlpBudd5BI-McLEjiRJsJNAGMVcu4Kh0pd91p1kIwBKo3kxKa-I8VZpACLWc6WaCgYKAWwSARESFQGOcNnC0gjR-O_-nqDHeezKB42gsQ0171",
	)
	ctx = metadata.NewOutgoingContext(ctx, headers)

	// cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	// conn, err := grpc.Dial("firestore.googleapis.com:443", cred)
	// errorc.CheckErr(err)
	// defer conn.Close()

	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	errorc.CheckErr(err)
	defer conn.Close()

	firestoreClient := firestore.NewFirestoreClient(conn)

	writeBidiStream, err := firestoreWrite.NewWrite(ctx, firestoreClient)
	errorc.CheckErr(err)

	writeBidiStream.Run()
}
