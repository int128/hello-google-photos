# Hello Google Photos Library API in Go

An example Go application using Google Photos Library API.

## TL;DR

1. Open https://console.cloud.google.com/apis/credentials
2. Create an OAuth client ID where the application type is other.
3. Set the following environment variables:

```sh
export GOOGLE_CLIENT_ID=
export GOOGLE_CLIENT_SECRET=
```

```
./hello-google-photos photo.jpg
```

The photo will be uploaded to your Google Photos Library.
