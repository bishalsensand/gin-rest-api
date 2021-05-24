package firebase

import "os"

var serviceAccountKeyfilePath = os.Getenv("SERVICE_ACCOUNT_KEY_FILE")
var projectId = os.Getenv("GCP_PROJECT_ID")
var collectionName = "entries"
