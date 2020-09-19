package testutil

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/sue445/primap/config"
	"google.golang.org/api/iterator"
)

// TestProjectID returns projectID for test
func TestProjectID() string {
	if config.GetProjectID() != "" {
		return config.GetProjectID()
	}

	return "test"
}

// CleanupFirestore cleanup Firestore data in test
func CleanupFirestore() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, TestProjectID())

	if err != nil {
		panic(err)
	}

	defer func() {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}()

	err = deleteCollection(ctx, client, client.Collection("Shops"), 100)
	if err != nil {
		panic(err)
	}
}

// https://firebase.google.com/docs/firestore/manage-data/delete-data#collections
func deleteCollection(ctx context.Context, client *firestore.Client, ref *firestore.CollectionRef, batchSize int) error {
	for {
		// Get a batch of documents
		iter := ref.Limit(batchSize).Documents(ctx)
		numDeleted := 0

		// Iterate through the documents, adding
		// a delete operation for each one to a
		// WriteBatch.
		batch := client.Batch()
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			batch.Delete(doc.Ref)
			numDeleted++
		}

		// If there are no documents to delete,
		// the process is over.
		if numDeleted == 0 {
			return nil
		}

		_, err := batch.Commit(ctx)
		if err != nil {
			return err
		}
	}
}
