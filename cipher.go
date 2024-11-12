package main

type Cipher interface {
	Encrypt(plaintext string, key interface{}) (string, error)
	Decrypt(ciphertext string, key interface{}) (string, error)
}
