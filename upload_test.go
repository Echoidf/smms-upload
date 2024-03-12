package main

import (
	"testing"
)

func TestUpload(t *testing.T) {
	token := "your token"
	base64Img := "..."
	url := "https://smms.app/api/v2/upload"

	res, err := uploadBase64(url, token, base64Img)
	if err != nil {
		t.Error("Error uploading file:", err)
		return
	}

	t.Log("Image URL:", res)
}
