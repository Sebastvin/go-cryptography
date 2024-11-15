package main

import (
	"fmt"
	"unicode"
)

type CaesarCipher struct{}

func NewCaesarCipher() *CaesarCipher {
	return &CaesarCipher{}
}

func (c *CaesarCipher) Encrypt(plaintext string, key interface{}) (string, error) {
	shift, ok := key.(int)

	if !ok {
		return "", fmt.Errorf("caesar cipher requires an integer key, got %T", key)
	}

	return shiftText(plaintext, shift)

}

func (c *CaesarCipher) Decrypt(plaintext string, key interface{}) (string, error) {
	shift, ok := key.(int)

	if !ok {
		return "", fmt.Errorf("caesar cipher requires an integer key, got %T", key)
	}

	return shiftText(plaintext, -shift)

}

func shiftText(text string, key int) (string, error) {
	encyptedText := ""
	switched := ""

	for _, char := range text {
		switched = getOffsetChar(char, key)
		encyptedText += switched
	}

	return encyptedText, nil
}

func getOffsetChar(character rune, offset int) string {
	if !unicode.IsLetter(character) {
		return string(character)
	}

	base := 'a'

	if unicode.IsUpper(character) {
		base = 'A'
	}

	offset = ((offset % 26) + 26) % 26
	shifted := rune((int(character-base) + offset) % 26)

	return string(shifted + base)
}
