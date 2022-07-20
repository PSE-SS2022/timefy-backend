package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"log"
)

const KEY_LOCATION = "AES_KEY_FILE"

var key = []byte{11, 108, 111, 57, 116, 83, 193, 127, 59, 57, 245, 188, 171, 59, 187, 101}

var IV = []byte("1234567812345678")

func createCipher() cipher.Block {
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return c
}

func Encrypt(plainText string) string {
	bytes := []byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return Encode(bytes)
}

func Decrypt(cipherText string) string {
	bytes := Decode(cipherText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return string(bytes)
}

func Encode(input []byte) string {
	return string(b64.StdEncoding.EncodeToString(input))
}

func Decode(input string) []byte {
	dcd, err := b64.StdEncoding.DecodeString(input)
	if err != nil {
		return []byte{}
	}
	return dcd
}
