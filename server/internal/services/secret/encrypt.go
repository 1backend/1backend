package secretservice

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// encrypt function encrypts the plaintext using the AES-256 algorithm
func encrypt(plaintext string, key string) (string, error) {
	// Ensure the key length is 32 bytes for AES-256
	if len(key) != 32 {
		return "", fmt.Errorf("key must be 32 bytes")
	}

	// Generate a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		return "", err
	}

	// Create AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Pad plaintext to be a multiple of the block size
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	paddedText := append([]byte(plaintext), bytes.Repeat([]byte{byte(padding)}, padding)...)

	// Encrypt the data
	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)

	// Return the IV and ciphertext as a base64 encoded string
	return base64.StdEncoding.EncodeToString(append(iv, ciphertext...)), nil
}

// decrypt function decrypts the ciphertext using the AES-256 algorithm
func decrypt(ciphertext string, key string) (string, error) {
	// Ensure the key length is 32 bytes for AES-256
	if len(key) != 32 {
		return "", fmt.Errorf("key must be 32 bytes")
	}

	// Decode the base64 encoded ciphertext
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Extract the IV from the first 16 bytes
	iv := data[:aes.BlockSize]
	ciphertextBytes := data[aes.BlockSize:]

	// Create AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Decrypt the data
	plaintext := make([]byte, len(ciphertextBytes))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertextBytes)

	// Remove padding
	padding := int(plaintext[len(plaintext)-1])
	plaintext = plaintext[:len(plaintext)-padding]

	return string(plaintext), nil
}
