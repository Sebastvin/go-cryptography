package main

import (
	"sort"
	"strings"
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

// https://www3.nd.edu/~busiforc/handouts/cryptography/letterfrequencies.html
var englishFrequency = map[rune]float64{
	'e': 11.1607, 'a': 8.4966, 'r': 7.5809, 'i': 7.5448,
	'o': 7.1635, 't': 6.9509, 'n': 6.6544, 's': 5.7351,
	'l': 5.4893, 'c': 4.5388, 'u': 3.6308, 'd': 3.3844,
	'p': 3.1671, 'm': 3.0129, 'h': 3.0034, 'g': 2.4705,
	'b': 2.0720, 'f': 1.8121, 'y': 1.7779, 'w': 1.2899,
	'k': 1.1016, 'v': 1.0074, 'x': 0.2902, 'z': 0.2722,
	'j': 0.1965, 'q': 0.1962,
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
		expectedFreq := englishFrequency[shiftedChar]
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
