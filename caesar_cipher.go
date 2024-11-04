package main
import (
	"unicode"
	"fmt"
)


func encrypt(plaintext string, key int) string {
	return shiftText(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return shiftText(ciphertext, -key)
}

func shiftText(text string, key int) string {
	encyptedText := ""
	switched := ""

	for _, char := range text {
		switched = getOffsetChar(char, key)
		encyptedText += switched
	}

	return encyptedText
}

func getOffsetChar(c rune, offset int) string {
	if unicode.IsLower(c) {
		return string(((c-'a'+rune(offset))%26+26)%26 + 'a')
	} else if unicode.IsUpper(c) {
		return string(((c-'A'+rune(offset))%26+26)%26 + 'A')
	}
	// Zwraca niezmieniony znak, jeśli to nie jest litera
	return string(c)
}

func main() {
	encrypted := encrypt("AbcdefgHIJ", 3)
	decrypted := decrypt(encrypted, 3)
	fmt.Println("Encrypted:", encrypted)	
	fmt.Println("Decrypted:", decrypted)	
}