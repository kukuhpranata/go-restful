package helper

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"strconv"
	"strings"
)

func Encrypt(id uint64) (string, error) {
	key := "Ik16vw84dbCxjTJGJwIFmUW4tfW8BT5u"
	value := int(id)
	plaintext := []byte(strconv.Itoa(value))

	key64, _ := base64.StdEncoding.DecodeString(key)
	// Define the cipher and key. Create a key size of 16 bytes (128 bits).
	cipherBlock, err := aes.NewCipher([]byte(key64[:16]))
	if err != nil {
		PanicIfError(err)
	}

	plainBytes := []byte(plaintext)
	blockSize := cipherBlock.BlockSize()

	// Padding the plaintext
	padding := blockSize - len(plainBytes)%blockSize
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	plainBytes = append(plainBytes, paddingText...)

	// Encrypt the plaintext
	encrypted := make([]byte, len(plainBytes))
	for bs, be := 0, blockSize; bs < len(plainBytes); bs, be = bs+blockSize, be+blockSize {
		cipherBlock.Encrypt(encrypted[bs:be], plainBytes[bs:be])
	}

	// Encode the result as base64
	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)

	// Replace '/' with '-' and '+' with '_'
	finalEncrypt := strings.NewReplacer("/", "-", "+", "_").Replace(encryptedBase64)

	return finalEncrypt, nil
}

func Decrypt(encryptedText string) (int, error) {
	key := "Ik16vw84dbCxjTJGJwIFmUW4tfW8BT5u"

	key64, _ := base64.StdEncoding.DecodeString(key)

	cipherBlock, err := aes.NewCipher([]byte(key64)[:16])
	if err != nil {
		PanicIfError(err)
	}

	// Replace '_' with '/' and '-' with '+'
	standardBase64 := strings.NewReplacer("_", "+", "-", "/").Replace(encryptedText)

	// Decode the base64 string
	encryptedBytes, err := base64.StdEncoding.DecodeString(standardBase64)
	if err != nil {
		PanicIfError(err)
	}

	decrypted := make([]byte, len(encryptedBytes))
	blockSize := cipherBlock.BlockSize()

	for bs, be := 0, blockSize; bs < len(encryptedBytes); bs, be = bs+blockSize, be+blockSize {
		cipherBlock.Decrypt(decrypted[bs:be], encryptedBytes[bs:be])
	}

	// Remove PKCS7 padding
	padding := decrypted[len(decrypted)-1]
	decrypted = decrypted[:len(decrypted)-int(padding)]

	valueDecrypted, err := strconv.Atoi(string(decrypted))
	if err != nil {
		PanicIfError(err)
	}
	return int(valueDecrypted), nil
}
