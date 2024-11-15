package main

import (
	"fmt"
	"sort"
	"strings"

	// "strings"
	"unicode"
)

type FrequencyPair struct {
	Letter rune
	Freq   float64
}

type DecryptResult struct {
	Text  string
	Shift int
	Score float64
}

// https://sjp.pwn.pl/poradnia/haslo/frekwencja-liter-w-polskich-tekstach;7072.html
var polishFrequency = map[rune]float64{
	'a': 8.91, 'i': 8.21, 'o': 7.75, 'e': 7.66, 'z': 5.64,
	'n': 5.52, 'r': 4.69, 'w': 4.65, 's': 4.32, 't': 3.98,
	'c': 3.96, 'y': 3.76, 'k': 3.51, 'd': 3.25, 'p': 3.13,
	'm': 2.80, 'u': 2.50, 'j': 2.28, 'l': 2.10, 'b': 1.47,
	'g': 1.42, 'h': 1.08, 'f': 0.30, 'q': 0.14, 'v': 0.04,
	'x': 0.02,
}

func calculateFrequency(text string) map[rune]float64 {
	freq := make(map[rune]float64)
	total := 0

	for _, char := range text {
		if unicode.IsLetter(char) {
			char = unicode.ToLower(char)
			freq[char]++
			total++
		}
	}

	for char := range freq {
		freq[char] = (freq[char] / float64(total)) * 100
	}
	return freq

}

func calculateDifference(textFreq map[rune]float64, shift int) float64 {
	difference := 0.0

	for char := 'a'; char <= 'z'; char++ {
		shiftedChar := rune((int(char-'a')+shift)%26 + int('a'))
		actualFreq := textFreq[char]
		expectedFreq := polishFrequency[shiftedChar]
		diff := actualFreq - expectedFreq
		difference += diff * diff
	}
	return difference
}

func decrypt(text string, shift int) string {
	result := strings.Builder{}

	for _, char := range text {
		if unicode.IsLetter(char) {
			isUpper := unicode.IsUpper(char)
			char = unicode.ToLower(char)

			decrypted := rune((int(char-'a')-shift+26)%26 + int('a'))

			if isUpper {
				decrypted = unicode.ToUpper(decrypted)
			}
			result.WriteRune(decrypted)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func breakCipher(ciphertext string) []DecryptResult {
	topN := 10

	textFreq := calculateFrequency(ciphertext)
	results := make([]DecryptResult, 0, 26)

	// SHIFT can be max 26
	for shift := 0; shift < 26; shift++ {
		score := calculateDifference(textFreq, shift)
		decrypted := decrypt(ciphertext, shift)
		results = append(results, DecryptResult{
			Text:  decrypted,
			Shift: shift,
			Score: score,
		})
	}

	// Sort by probality
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score < results[0].Score
	})

	return results[:topN]
}

func main() {
	ciphertext := "olssv dvysk"

	results := breakCipher(ciphertext)

	for i, result := range results {
		fmt.Printf("--------------- Result  %d -----------------\n", i+1)
		fmt.Printf("Shifts: %d\n", result.Shift)
		fmt.Printf("Text: %s\n", result.Text)
		fmt.Printf("Matching rate: %.2f\n", result.Score)
	}
}
