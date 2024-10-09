package tools

import (
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {

	// 32 bytes
	secretKey := []byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc47")
	// This will successfully encrypt & decrypt
	ciphertext1 := AesEncrypt(secretKey, "This is some sensitive information")
	fmt.Printf("Encrypted ciphertext 1: %x \n", ciphertext1)

	plaintext1 := AesDecrypt(secretKey, ciphertext1)
	fmt.Printf("Decrypted plaintext 1: %s \n", plaintext1)

	// This will successfully encrypt & decrypt as well.
	ciphertext2 := AesEncrypt(secretKey, "Hello")
	fmt.Printf("Encrypted ciphertext 2: %x \n", ciphertext2)

	plaintext2 := AesDecrypt(secretKey, ciphertext2)
	fmt.Printf("Decrypted plaintext 2: %s \n", plaintext2)
}
