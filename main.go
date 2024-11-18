package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		fmt.Printf("--------------- Cryptography -----------------\n\n")
		fmt.Printf("Choose one of action below: \n")
		fmt.Printf("1 - Caesar Cipher \n")
		fmt.Printf("2 - Vigenere cipher \n")
		fmt.Printf("3 - Caesar breaker \n")

		var choice int
		in := bufio.NewReader(os.Stdin)

		fmt.Printf("Choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("--------------- Caesar Cipher -----------------\n\n")
			var caesarCipher = NewCaesarCipher()

			fmt.Printf("Type plaintext: ")
			plaintext, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Type key: ")
			keyStr, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			keyStr = strings.TrimSpace(keyStr)
			key, err := strconv.Atoi(keyStr)

			if err != nil {
				fmt.Println(err)
			}

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
		case 2:
			fmt.Printf("--------------- Vigenere Cipher -----------------\n\n")

			fmt.Printf("Type plaintext: ")
			plaintext, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Type key: ")
			keyStr, err := in.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			keyStr = strings.TrimSpace(keyStr)

			result, err := generateKey(plaintext, keyStr)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Generated key: %s\n", result)

			var vigenereCipher = NewVigenereCipher()
			encrypted, err := vigenereCipher.Encrypt(plaintext, result)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Encrypted: %s\n", encrypted)
			decrypted, err := vigenereCipher.Decrypt(encrypted, result)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Decrypted: %s\n", decrypted)

		case 3:
			fmt.Printf("--------------- Caesar breaker -----------------\n\n")

			fmt.Print("Type your text to break: ")
			cipheredText, err := in.ReadString('\n')

			if err != nil {
				fmt.Println(err)
			}

			results := breakCipher(cipheredText)

			for i, result := range results {
				fmt.Printf("--------------- Result  %d -----------------\n", i+1)
				fmt.Printf("Shifts: %d\n", result.Shift)
				fmt.Printf("Text: %s\n", result.Text)
				fmt.Printf("Matching rate: %.2f\n", result.Score)
			}
		default:
			fmt.Println("Invalid option. Please choose correct one")
		}

		fmt.Println("Continue? Y/N")
		answer, err := in.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToUpper(answer)

		if err != nil {
			fmt.Println(err)
		}

		switch answer {
		case "Y", "YES":
			continue
		default:
			return
		}
	}
}
