package firebase

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createClient(ctx context.Context) *firestore.Client {
	var options []option.ClientOption

	if serviceAccountKeyfilePath != "" {
		// When running locally, service accounst key file is used for authenticating with Firestore
		// Authentication is automatic when running in GCP
		options = append(options, option.WithCredentialsFile(serviceAccountKeyfilePath))
	}

	client, clientErr := firestore.NewClient(ctx, projectId, options...)

	if clientErr != nil {
		log.Fatalf("Failed to create client: %v", clientErr)
	}

	return client
}

func postEntryToFirebase(entry Entry) (string, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docRef, _, err := client.
		Collection(collectionName).
		Add(ctx, entry)

	if err != nil {
		return "", err
	}

	return docRef.ID, nil
}

func getEntryFromFirebase(entryId string) (*Entry, error) {
	if entryId == "" {
		return nil, fmt.Errorf("Entry id is empty")
	}
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	docSnapshot, err := client.
		Collection(collectionName).
		Doc(entryId).
		Get(ctx)

	if err != nil {

		// Not found
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}

		return nil, err
	}

	if docSnapshot.Exists() != true {
		return nil, nil
	}

	var entry Entry
	docSnapshot.DataTo(&entry)
	entry.Id = &docSnapshot.Ref.ID
	return &entry, nil
}

func getAllEntriesFromFirebase(count int) ([]Entry, error) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	iter := client.
		Collection(collectionName).
		Limit(count).
		Documents(ctx)

	var fetchedEntries []Entry = []Entry{}

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("Error occured when reading firestore iterator %v", err.Error())
		}

		var entry Entry

		_ = doc.DataTo(&entry)
		entry.Id = &doc.Ref.ID
		fetchedEntries = append(fetchedEntries, entry)
	}
	return fetchedEntries, nil
}

func deleteFirebaseEntryById(entryId string) (bool, error) {

	if entryId == "" {
		return false, fmt.Errorf("Entry id is empty")
	}

	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, err := client.
		Collection(collectionName).
		Doc(entryId).
		Delete(ctx)

	return (err == nil), err
}

func updateFirebaseEntryById(entryId string, entry Entry) (bool, error) {
	if entryId == "" {
		return false, fmt.Errorf("Entry id is empty")
	}

	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, setErr := client.
		Collection(collectionName).
		Doc(entryId).
		Set(ctx, entry)

	return (setErr == nil), setErr

}
