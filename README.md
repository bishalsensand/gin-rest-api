# Quick and dirty REST API using gin
This is an example of quick and incomplete REST API written in [Gin Web Framework](https://github.com/gin-gonic/gin)


## Configuring
See `firebase/config.go` for configuration options
## Running
Run `go run .` to start a webserver in port `8080`. Routes defined in `main.go`.


## Deploying
To deploy to Google Cloud, make sure you have a project that you have correct access to, and run these commands

```bash
# Set env vars
export GCP_PROJECT_ID=bishal-playground-313602
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