package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const apiVersion = "v1"
const basePath = "https://photoslibrary.googleapis.com/"

// PhotosHelper is a client for uploading a media.
// photoslibrary does not provide `/v1/uploads` API so we implement here.
type PhotosHelper struct {
	client *http.Client
}

// NewPhotosHelper creates a new client.
func NewPhotosHelper(client *http.Client) *PhotosHelper {
	return &PhotosHelper{client}
}

// Upload sends the media and returns the UploadToken.
func (c *PhotosHelper) Upload(r io.Reader, filename string) (string, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/uploads", basePath, apiVersion), r)
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Goog-Upload-File-Name", filename)

	res, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	uploadToken := string(b)
	return uploadToken, nil
}
