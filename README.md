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
./hello-google-photos photo1.jpg photo2.jpg
```

The photos will be uploaded to your Google Photos Library.

## How it works

This depends on [google/google-api-go-client](https://github.com/google/google-api-go-client).

It does not provide [media uploads](https://developers.google.com/photos/library/guides/upload-media#uploading-bytes) and this implements it using `http.Client`.

## Caveats

This does not persist an access token.
You must do browser authentication every time.

This supports does not support a service account.
