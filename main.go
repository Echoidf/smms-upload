package main

import (
	"fmt"
)

func main() {
	filePath := "./codearena.png"

	res, err := uploadLocalFile(filePath)
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	fmt.Println("Image URL:", res)
}
