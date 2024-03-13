package main

import (
	"testing"
)

func TestUpload(t *testing.T) {
	base64Img := "..."

	res, err := uploadBase64(base64Img)
	if err != nil {
		t.Error("Error uploading file:", err)
		return
	}

	t.Log("Image URL:", res)
}
