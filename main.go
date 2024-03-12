package main

import (
	"fmt"
)

func main() {
	token := "your token"
	filePath := "./codearena.png"
	url := "https://smms.app/api/v2/upload"

	res, err := uploadFile(url, token, filePath)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	fmt.Println("Image URL:", res)
}
