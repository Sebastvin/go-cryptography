package main
import (
	"strings"
	"fmt"
)


func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	encypted_str := ""
	switched := ""

	for _, char := range text {
		switched = getOffsetChar(char, key)
		encypted_str += switched
	}

	return encypted_str
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	alphabetLength := len(alphabet)
	
	i := strings.Index(alphabet, string(c))

	if i == -1 {
		return ""
	}
	

	newIndex := (i + offset) % alphabetLength
	
	if newIndex < 0 {
		newIndex += alphabetLength
	}

	return string(alphabet[newIndex])
}

func main() {
	// Test cases
	// 	{"abcdefghi", 1, "bcdefghij"},
	// 	{"hello", 5, "mjqqt"},
	// 	{"correcthorsebatterystaple", 16, "sehhusjxehiurqjjuhoijqfbu"},
	// 	{"onetwothreefourfivesixseveneightnineten", 25, "nmdsvnsgqddentqehudrhwrdudmdhfgsmhmdsdm"},

	encrypted := encrypt("abcdefghi", 1)
	decrypted := decrypt(encrypted, 1)
	fmt.Println(encrypted)	
	fmt.Println(decrypted)	
}