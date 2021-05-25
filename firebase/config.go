package firebase

import "os"

// Path to service-account-key.json file downloaded from firebase
var serviceAccountKeyfilePath = os.Getenv("SERVICE_ACCOUNT_KEY_FILE")

// Your google cloud project ID
var projectId = os.Getenv("GCP_PROJECT_ID")

// Name of the collection to use in firestore
var collectionName = "entries"
