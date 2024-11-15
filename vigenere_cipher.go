package main

import (
	"fmt"
	"unicode"
)

type VigenereCipher struct{}

func NewVigenereCipher() *VigenereCipher {
	return &VigenereCipher{}
}

func (v *VigenereCipher) Encrypt(plaintext string, key interface{}) (string, error) {

	shift, ok := key.(string)

	if !ok {
		return "", fmt.Errorf("vignere cipher requires an string key, got %T", key)
	}

	for _, char := range shift {
		if !unicode.IsLetter(char) {
			return "", fmt.Errorf("key must contain only letters, found: %c", char)
		}
	}

	if len(shift) == 0 {
		return "", fmt.Errorf("key cannot be empty")
	}
	cipherText := make([]rune, len(plaintext))
	keyIndex := 0

	for i, char := range plaintext {
		//  Skip whitespace chars
		if !unicode.IsLetter(char) {
			cipherText[i] = char
			continue
		}

		isUpper := unicode.IsUpper(char)

		// Convert to 0-25 range
		var base rune
		if isUpper {
			base = 'A'
		} else {
			base = 'a'
		}

		keyChar := unicode.ToUpper(rune(shift[keyIndex%len(shift)])) - 'A'

		plainChar := char - base
		encrypted := (plainChar + keyChar) % 26
		cipherText[i] = encrypted + base

		keyIndex++
	}

	return string(cipherText), nil
}

func (v *VigenereCipher) Decrypt(ciphertext string, key interface{}) (string, error) {

	shift, ok := key.(string)

	if !ok {
		return "", fmt.Errorf("vignere cipher requires an string key, got %T", key)
	}

	for _, char := range shift {
		if !unicode.IsLetter(char) {
			return "", fmt.Errorf("key must contain only letters, found: %c", char)
		}
	}

	if len(shift) == 0 {
		return "", fmt.Errorf("key cannot be empty")
	}

	originalText := ""
	keyIndex := 0

	for i := 0; i < len(ciphertext); i++ {
		char := ciphertext[i]

		if !unicode.IsLetter(rune(char)) {
			originalText += string(char)
			continue
		}

		isUpper := unicode.IsUpper(rune(char))
		var base byte
		if isUpper {
			base = 'A'
		} else {
			base = 'a'
		}

		keyChar := unicode.ToUpper(rune(shift[keyIndex%len(shift)])) - 'A'

		charValue := byte(char) - base
		decrypted := byte((int(charValue) - int(keyChar) + 26) % 26)
		originalText += string(decrypted + base)

		keyIndex++
	}

	return originalText, nil
}

func generateKey(plaintext string, key string) (string, error) {
	// TODO: Change variable naming
	x := len(plaintext)
	i := 0

	for {
		if x == i {
			i = 0
		}

		if len(key) == len(plaintext) {
			break
		}
		key += string(key[i])
		i++
	}

	return key, nil
}
