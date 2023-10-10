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
		"Authorization", "Bearer "+"ya29.a0AfB_byCSAhjXmMtLs9uLU1A9zj5l5sBbK0DLuV7o0BTKpHpnF2xEJqYn4k3I5Ofcusmv5e7t1ZAd48HLKdVru0CwQmd4QzaYpqlMgi93SpeYEPGaPBWhcD44bPKfCkbFnIN4w1M9ysthX2-V6IWq0c-ju-H8I-j4mwhfaCgYKAeMSARESFQGOcNnCSEiePD7PKYz5l3-qMKoIdg0171",
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
