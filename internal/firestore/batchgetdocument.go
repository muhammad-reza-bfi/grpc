package firestoreclient

import (
	"context"
	"errors"
	"fmt"
	"io"

	firestore "google.golang.org/genproto/googleapis/firestore/v1beta1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BatchGetDocument struct {
	BatchGetDocumentClient firestore.Firestore_BatchGetDocumentsClient
}

func NewBatchGetDocument(ctx context.Context, firestoreClient firestore.FirestoreClient) (*BatchGetDocument, error) {
	writeClient, err := firestoreClient.BatchGetDocuments(ctx, &firestore.BatchGetDocumentsRequest{
		Database: "projects/elangreza-golang-base/databases/(default)",
		Documents: []string{
			"projects/elangreza-golang-base/databases/(default)/documents/a/test",
			"projects/elangreza-golang-base/databases/(default)/documents/a/test send",
		},
	})
	if err != nil {
		return nil, err
	}

	return &BatchGetDocument{
		BatchGetDocumentClient: writeClient,
	}, nil
}

func (w *BatchGetDocument) Run() error {

	for {
		res, err := w.BatchGetDocumentClient.Recv()

		if sts, ok := status.FromError(err); ok && sts.Code() == codes.Canceled {
			fmt.Printf("got err 1 %v \n", sts.Code())
			return fmt.Errorf("got error %v", sts.Code())
		} else if err == io.EOF {
			fmt.Printf("got err 2 %v \n", err.Error())
			return errors.New("stream closed")
		} else if err != nil {
			fmt.Printf("got err 3 %v \n", err.Error())
			return err
		}

		fmt.Printf("got value: %v \n", res.GetFound().GetFields())
	}
}
