package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	var text = "ini adalah text rahasia"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
}
