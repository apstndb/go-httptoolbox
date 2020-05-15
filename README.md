# HTTP Toolbox

HTTP handlers for digging serverless services.

## Deploy

There are examples for App Engine, Cloud Functions, Cloud Run and [GCP Buildpacks](https://github.com/GoogleCloudPlatform/buildpacks).

The handlers can be security backdoor.
Don't deploy without some security setting.

### App Engine Go 1.11

```
$ gcloud app deploy app_go111.yaml 
```

### App Engine Go 1.12

```
$ gcloud app deploy app_go112.yaml 
```

### Cloud Run

```
# Use Cloud Build
$ gcloud builds submit -t gcr.io/${PROJECT_ID}/httptoolbox 
# or use GCP Buildpacks
$ pack build --builder gcr.io/buildpacks/builder -e GOOGLE_BUILDABLE=./cmd/main --publish gcr.io/${PROJECT_ID}/httptoolbox 
$ gcloud run deploy --image gcr.io/${PROJECT_ID}/httptoolbox 
```

```
$ gcloud run deploy --image gcr.io/${PROJECT_ID}/httptoolbox 
```

### Cloud Functions

```
$ gcloud functions deploy ExecDmesg --trigger-http --runtime=go111
```

### GCP Buildpacks for App

```
$ pack build --builder gcr.io/buildpacks/builder -e GOOGLE_BUILDABLE=./cmd/main go-httptoolbox-app
$ docker run -p 8080:8080 --rm go-httptoolbox-app 
```

You can use `gcloud alpha builds submit --pack`.

```
$ gcloud alpha builds submit --pack=builder=gcr.io/buildpacks/builder,env=GOOGLE_BUILDABLE=./cmd/main,image=gcr.io/${PROJECT_ID}/go-httptoolbox-app
```

### GCP Buildpacks for Function

```
$ pack build --builder gcr.io/buildpacks/builder -e GOOGLE_FUNCTION_TARGET=DumpRequest go-httptoolbox-function

$ docker run -p 8080:8080 --rm go-httptoolbox-function 
```

You can use `gcloud alpha builds submit --pack`.

```
$ gcloud alpha builds submit --pack=builder=gcr.io/buildpacks/builder,env=GOOGLE_FUNCTION_TARGET=DumpRequest,image=gcr.io/${PROJECT_ID}/go-httptoolbox-function
```
