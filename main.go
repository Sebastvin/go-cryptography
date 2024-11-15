package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("--------------- Caesar Cipher -----------------\n\n")
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

	fmt.Printf("--------------- Vigenere Cipher -----------------\n\n")

	plaintext = "kanguaro"
	key_str := "KEY"

	result, err := generateKey(plaintext, key_str)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated key: %s\n", result)

	var vigenereCipher = NewVigenereCipher()
	encrypted, err := vigenereCipher.Encrypt(plaintext, result)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypted: %s\n", encrypted)
	decrypted, err := vigenereCipher.Decrypt(encrypted, result)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)

	fmt.Printf("--------------- Caesar breaker -----------------\n\n")
	ciphertext := "olssv dvysk"

	results := breakCipher(ciphertext)

	for i, result := range results {
		fmt.Printf("--------------- Result  %d -----------------\n", i+1)
		fmt.Printf("Shifts: %d\n", result.Shift)
		fmt.Printf("Text: %s\n", result.Text)
		fmt.Printf("Matching rate: %.2f\n", result.Score)
	}
}
