package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	menuTemplate = `
--------------- Cryptography -----------------

Choose one of action below:
1 - Caesar Cipher
2 - Vigenere Cipher
3 - Caesar Breaker

Choice: `

	caesarHeader = `
--------------- Caesar Cipher -----------------
`
	vigenereHeader = `
--------------- Vigenere Cipher -----------------
`
	breakerHeader = `
--------------- Caesar Breaker -----------------
`
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		choice, err := getMenuChoice()

		if err != nil {
			log.Printf("error getting choice: %v", err)
			continue
		}

		switch choice {
		case 1:
			handleCaesarCipher(reader)
		case 2:
			handleVigenereCipher(reader)
		case 3:
			handleCaesarBreaker(reader)
		default:
			fmt.Println("Invalid option. Please choose correct one")
			continue
		}

		if !shouldContinue(reader) {
			return
		}
	}
}

func getMenuChoice() (int, error) {
	fmt.Print(menuTemplate)
	var choice int
	_, err := fmt.Scanln(&choice)
	return choice, err
}

func handleCaesarCipher(reader *bufio.Reader) {
	fmt.Print(caesarHeader)
	caesarCipher := NewCaesarCipher()

	plaintext, err := readInput(reader, "Type plaintext: ")
	if err != nil {
		log.Printf("Error reading plaintext: %v", err)
		return
	}

	keyStr, err := readInput(reader, "Type key: ")
	if err != nil {
		log.Printf("Error reading key: %v", err)
		return
	}

	key, err := strconv.Atoi(strings.TrimSpace(keyStr))
	if err != nil {
		log.Printf("Invalid key format: %v", err)
		return
	}

	encrypted, err := caesarCipher.Encrypt(plaintext, key)
	if err != nil {
		log.Printf("Caesar encryption error: %v", err)
		return
	}

	decrypted, err := caesarCipher.Decrypt(encrypted, key)
	if err != nil {
		log.Printf("Caesar decryption error: %v", err)
		return
	}

	fmt.Println("Encrypted:", encrypted)
	fmt.Println("Decrypted:", decrypted)
}

func handleVigenereCipher(reader *bufio.Reader) {
	fmt.Print(vigenereHeader)

	plaintext, err := readInput(reader, "Type plaintext: ")
	if err != nil {
		log.Printf("Error reading plaintext: %v", err)
		return
	}

	keyStr, err := readInput(reader, "Type key: ")
	if err != nil {
		log.Printf("Error reading key: %v", err)
		return
	}

	key, err := generateKey(plaintext, strings.TrimSpace(keyStr))
	if err != nil {
		log.Printf("Error generating key: %v", err)
		return
	}

	vigenereCipher := NewVigenereCipher()
	encrypted, err := vigenereCipher.Encrypt(plaintext, key)
	if err != nil {
		log.Printf("Vigenere encryption error: %v", err)
		return
	}

	decrypted, err := vigenereCipher.Decrypt(encrypted, key)
	if err != nil {
		log.Printf("Vigenere decryption error: %v", err)
		return
	}

	// fmt.Println("Generated key:", key)
	fmt.Println("Encrypted:", encrypted)
	fmt.Println("Decrypted:", decrypted)
}

func handleCaesarBreaker(reader *bufio.Reader) {
	fmt.Print(breakerHeader)

	ciphertext, err := readInput(reader, "Type your text to break: ")
	if err != nil {
		log.Printf("Error reading ciphertext: %v", err)
		return
	}

	topNStr, err := readInput(reader, "Type your topN: ")
	if err != nil {
		log.Printf("Error reading input: %v", err)
		return
	}

	topNStr = strings.Replace(topNStr, "\n", "", 1)
	topNStr = strings.TrimSpace(topNStr)

	topN, err := strconv.Atoi(topNStr)
	if err != nil {
		log.Printf("Error converting topN to integer: %v", err)
		return
	}

	results := breakCipher(ciphertext, topN)
	for i, result := range results {
		fmt.Printf("--------------- Result %d -----------------\n", i+1)
		fmt.Printf("Shifts: %d\n", result.Shift)
		fmt.Printf("Text: %s\n", result.Text)
		fmt.Printf("Matching rate: %.2f\n", result.Score)
	}
}

func readInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	return reader.ReadString('\n')
}

func shouldContinue(reader *bufio.Reader) bool {
	fmt.Print("Continue? (Y/N): ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading continue choice: %v", err)
		return false
	}

	answer = strings.ToUpper(strings.TrimSpace(answer))
	return answer == "Y" || answer == "YES"
}
