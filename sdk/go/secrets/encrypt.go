package secrets

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// encrypt function encrypts the plaintext using the AES-256 algorithm
func Encrypt(plaintext string, key string) (string, error) {
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
func Decrypt(ciphertext string, key string) (string, error) {
	// Ensure the key length is 32 bytes for AES-256
	if len(key) != 32 {
		return "", fmt.Errorf("key must be 32 bytes")
	}

	// Decode the base64 encoded ciphertext
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Need at least one block for IV
	if len(data) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	ciphertextBytes := data[aes.BlockSize:]

	// Ciphertext must be non-empty and block aligned
	if len(ciphertextBytes) == 0 || len(ciphertextBytes)%aes.BlockSize != 0 {
		return "", fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	// Create AES cipher block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Decrypt the data
	plaintext := make([]byte, len(ciphertextBytes))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertextBytes)

	// Validate and remove PKCS#7 padding
	if len(plaintext) == 0 {
		return "", fmt.Errorf("decryption produced empty plaintext")
	}
	padding := int(plaintext[len(plaintext)-1])
	if padding == 0 || padding > aes.BlockSize || padding > len(plaintext) {
		return "", fmt.Errorf("invalid padding size")
	}

	// Check all padding bytes
	for i := 0; i < padding; i++ {
		if plaintext[len(plaintext)-1-i] != byte(padding) {
			return "", fmt.Errorf("invalid padding bytes")
		}
	}
	plaintext = plaintext[:len(plaintext)-padding]

	return string(plaintext), nil
}
