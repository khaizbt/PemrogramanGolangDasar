package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "Khaiz Badaru Tammam"
	var encodeString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodeString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodeString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

}
