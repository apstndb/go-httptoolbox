# HTTP Toolbox

HTTP handlers for digging serverless services.


## Deploy

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
$ gcloud builds submit -t gcr.io/${PROJECT_ID}/httptoolbox 
$ gcloud run deploy --image gcr.io/${PROJECT_ID}/httptoolbox 
```

### Cloud Functions

```
$ gcloud functions deploy ExecDmesg --trigger-http --runtime=go111
```

