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

func getOffsetChar(c rune, offset int) string {
	if unicode.IsLower(c) {
		return string(((c-'a'+rune(offset))%26+26)%26 + 'a')
	} else if unicode.IsUpper(c) {
		return string(((c-'A'+rune(offset))%26+26)%26 + 'A')
	}

	return string(c)
}
