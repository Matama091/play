package layers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

func padByPkcs7(data []byte) []byte {
	padSize := aes.BlockSize
	if len(data)%aes.BlockSize != 0 {
		padSize = aes.BlockSize - (len(data))%aes.BlockSize
	}

	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, pad...)
}

func unPadByPkcs7(data []byte) []byte {
	padSize := int(data[len(data)-1])
	return data[:len(data)-padSize]
}

// EncryptPacket encrypts the packet.
func EncryptPacket(buf, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	paddBuf := padByPkcs7(buf)
	cipherBuf := make([]byte, aes.BlockSize+len(paddBuf))
	iv := cipherBuf[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		fmt.Printf("err: %s\n", err)
	}

	encryptStream := cipher.NewCBCEncrypter(block, iv)
	encryptStream.CryptBlocks(cipherBuf[aes.BlockSize:], paddBuf)

	return cipherBuf
}

// DecryptPacket decrypts the packet.
func DecryptPacket(buf, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("err: %s\n", err)
	}

	decryptedText := make([]byte, len(buf[aes.BlockSize:]))
	decryptStream := cipher.NewCBCDecrypter(block, buf[:aes.BlockSize])
	decryptStream.CryptBlocks(decryptedText, buf[aes.BlockSize:])

	return unPadByPkcs7(decryptedText)
}
