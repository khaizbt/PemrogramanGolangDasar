package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("/////////// Penerapan Fungsi Encode dan Decode ////////")

	var data1 = "Khafif Damanghuri"

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data1)))
	base64.StdEncoding.Encode(encoded, []byte(data1))
	var encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)

	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Println(decoded)
	var decodedString = string(decoded)
	fmt.Println(decodedString)
}
