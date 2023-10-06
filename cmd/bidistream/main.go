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
		"Authorization", "Bearer "+"ya29.a0AfB_byCEnEWVfwupqwZuw_ykib3QrsZvFu_RTfN26BzWSaZc_4V_rbG5rleM1x8SDsbsYb2kUIvbTDBwCYe7HDpD_p_lv7TZD_E6Y_LXvocqSZaR7wiAaV67UIObeyE3HI1RLHv6t05fetWCiC3qzJUkSkFFHdXUF0mgaCgYKASoSARESFQGOcNnCKXTcB7ZK2G6Ya92E4BNCUg0171",
	)
	ctx = metadata.NewOutgoingContext(ctx, headers)

	cred := grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))
	conn, err := grpc.Dial("firestore.googleapis.com:443", cred)
	errorc.CheckErr(err)
	defer conn.Close()

	firestoreClient := firestore.NewFirestoreClient(conn)

	writeBidiStream, err := firestoreclient.NewWrite(ctx, firestoreClient)
	errorc.CheckErr(err)

	writeBidiStream.Run()
}
