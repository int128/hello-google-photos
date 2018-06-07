package main

import (
	"context"
	"log"
	"os"
	"path"

	photoslibrary "google.golang.org/api/photoslibrary/v1"
)

func main() {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal(`Error: GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET must be set.
1. Open https://console.cloud.google.com/apis/credentials
2. Create an OAuth client ID where the application type is other.
3. Set the following environment variables:
export GOOGLE_CLIENT_ID=
export GOOGLE_CLIENT_SECRET=
`)
	}

	filepath := os.Args[1]
	filename := path.Base(filepath)

	ctx := context.Background()
	client, err := NewOAuthClient(ctx, clientID, clientSecret)
	if err != nil {
		log.Fatal(err)
	}

	helper := NewPhotosHelper(client)
	log.Printf("Uploading %s", filename)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	uploadToken, err := helper.Upload(file, filename)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Uploaded %s as token %s", filename, uploadToken)

	photos, err := photoslibrary.New(client)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Adding media %s", filename)
	batch, err := photos.MediaItems.BatchCreate(&photoslibrary.BatchCreateMediaItemsRequest{
		NewMediaItems: []*photoslibrary.NewMediaItem{
			&photoslibrary.NewMediaItem{
				Description:     filename,
				SimpleMediaItem: &photoslibrary.SimpleMediaItem{UploadToken: uploadToken},
			},
		},
	}).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range batch.NewMediaItemResults {
		log.Printf("Added media %s as %s", result.MediaItem.Description, result.MediaItem.Id)
	}
}
