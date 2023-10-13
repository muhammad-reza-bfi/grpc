package main

import (
	"context"
	"flag"
	"fmt"

	errorc "github.com/elangreza14/grpc/internal/error"
	firestoreWrite "github.com/elangreza14/grpc/internal/firestore/write"
	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// gen token from here
// https://developers.google.com/oauthplayground/?code=4/0AfJohXkE9e9rW6U7oYpAi5f4khtyexclBV8wr6QZImZtRq8IoPv_Q2MuClbhXy8SYbAhwg&scope=https://www.googleapis.com/auth/cloud-platform

func main() {

	forkPtr := flag.Bool("mock", false, "mock server")
	flag.Parse()
	ctx := context.Background()

	headers := metadata.Pairs(
		"Authorization", "Bearer "+"ya29.a0AfB_byA7yn9Du-Dj6l47wMB-2AYHjBEcK2hgN5BH5ESue4xW-RWFdUL51hxVfKY_jVt4eRlEnVsrUoY_Mh1uGqTRXlJgg46242CQbaYilf15c7XYvhSsxOiEIjGSved3Uz5SbvNK92AgN6VucDg5T-fOqFvizdQWf1kFaCgYKAZcSARESFQGOcNnCLj9Gfj9DDjPWGxYi3EV0cQ0171",
	)
	ctx = metadata.NewOutgoingContext(ctx, headers)

	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))

	var err error
	var conn *grpc.ClientConn

	if *forkPtr {
		fmt.Println("running mock server")
		conn, err = grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		fmt.Println("running firebase server")
		conn, err = grpc.Dial("firestore.googleapis.com:443", cred)
	}
	errorc.CheckErr(err)
	defer conn.Close()

	firestoreClient := firestore.NewFirestoreClient(conn)

	writeBidiStream, err := firestoreWrite.NewWrite(ctx, firestoreClient)
	errorc.CheckErr(err)

	writeBidiStream.Run()
}
