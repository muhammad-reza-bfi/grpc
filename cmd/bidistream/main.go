package main

import (
	"context"

	errorc "github.com/elangreza14/grpc/internal/error"
	firestoreWrite "github.com/elangreza14/grpc/internal/firestore/write"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// gen token from here
// https://developers.google.com/oauthplayground/?code=4/0AfJohXkE9e9rW6U7oYpAi5f4khtyexclBV8wr6QZImZtRq8IoPv_Q2MuClbhXy8SYbAhwg&scope=https://www.googleapis.com/auth/cloud-platform

func main() {
	ctx := context.Background()

	headers := metadata.Pairs(
		"Authorization", "Bearer "+"ya29.a0AfB_byAiJ5I2cRnppWJ2FEGp5Af9PqlCGUPqIKSJwa8AH8rvP0CkU14zNWQgg8KNTwMJJTn6TXbyupX13PgNE0YrzmsuDSmC5adTrivrYVkESafqzQAgjqxOvAEkIfJ8iu-Sqy230tqfEOvvRDVNt4iie3z3jSvv3N3laCgYKAZQSARESFQGOcNnCDkrA6Rk4E4ACa9c0af2L5g0171",
	)
	ctx = metadata.NewOutgoingContext(ctx, headers)

	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial("firestore.googleapis.com:443", cred)
	errorc.CheckErr(err)
	defer conn.Close()

	// conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// errorc.CheckErr(err)
	// defer conn.Close()

	firestoreClient := firestore.NewFirestoreClient(conn)

	writeBidiStream, err := firestoreWrite.NewWrite(ctx, firestoreClient)
	errorc.CheckErr(err)

	writeBidiStream.Run()
}
