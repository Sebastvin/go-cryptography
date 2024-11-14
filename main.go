package main

import (
	"fmt"
	"log"
)

func main() {

	var caesarCipher = NewCaesarCipher()

	plaintext := "HIROSHIMA NAGASAKI"
	key := 3

	caesarEncrypted, err := caesarCipher.Encrypt(plaintext, key)
	if err != nil {
		log.Fatal("Caesar encryption error:", err)
	}

	caesarDecrypted, err := caesarCipher.Decrypt(caesarEncrypted, key)
	if err != nil {
		log.Fatal("Caesar decryption error:", err)
	}

	fmt.Printf("Encrypted: %s\n", caesarEncrypted)
	fmt.Printf("Decrypted: %s\n", caesarDecrypted)
}
