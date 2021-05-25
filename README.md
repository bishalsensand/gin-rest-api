# Quick and dirty Firestore REST API using gin
This is an example of quick and incomplete REST API written in [Gin Web Framework](https://github.com/gin-gonic/gin). The project uses [Google Cloud's Firestore](https://cloud.google.com/firestore) for data persistence.


## Configuring
Access to firestore database is provided by using 2 environment variables. 

- `SERVICE_ACCOUNT_KEY_FILE` is path to a json key file downloaded from firebase project. This variable is not required to be set when code is executing in Google cloud compute environment.
- `GCP_PROJECT_ID` is the name of google cloud project the firestore databse lives in. This is a required config.

See `firebase/config.go` for configuration options

## Running
Run `go run .` to start a webserver in port `8080`. Routes defined in `main.go`.


## Deploying
To deploy to Google Cloud, make sure you have a project that you have correct access to, and run these commands

```bash
# Set env vars
export GCP_PROJECT_ID=your-project-name
export APP_BUILD_ID=gin-rest-api

# build containerized image and push to google container registry
gcloud builds submit --tag gcr.io/${GCP_PROJECT_ID}/${APP_BUILD_ID}

# deploy image to google cloud run
gcloud run deploy \
    --image gcr.io/${GCP_PROJECT_ID}/${APP_BUILD_ID} \
    --platform managed \
    --max-instances 1 \
    --platform managed \
    --port 8080 \
    --timeout 10s \
    --allow-unauthenticated \
    --region australia-southeast1\
    ${APP_BUILD_ID}
```