package main

import (
	"bytes"
	"crypto/cipher"

	"github.com/codahale/aesnicheck"
)

func pad(data []byte, blockSize int) []byte {
	padSize := blockSize - len(data)%blockSize
	ret := make([]byte, len(data)+padSize)
	copy(ret, data)
	copy(ret[len(data):], bytes.Repeat([]byte{byte(padSize)}, padSize))
	return ret
}

// func unpad(data []byte) []byte {
// 	padSize := int(data[len(data)-1])
// 	if padSize > len(data) {
// 		return nil
// 	}
// 	return data[:len(data)-padSize]
// }

func Check() bool {
	return aesnicheck.HasAESNI()
}

func EncryptWithCBC(block cipher.Block, iv []byte, data []byte) []byte {
	data = pad(data, block.BlockSize())
	enc := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(data))
	enc.CryptBlocks(encrypted, data)
	return encrypted
}

func DecryptWithCBC(block cipher.Block, iv []byte, data []byte) []byte {
	dec := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	dec.CryptBlocks(decrypted, data)
	return decrypted
}

func EncryptWithCFB(block cipher.Block, iv []byte, data []byte) []byte {
	enc := cipher.NewCFBEncrypter(block, iv)
	encrypted := make([]byte, len(data))
	enc.XORKeyStream(encrypted, data)
	return encrypted
}

func DecryptWithCFB(block cipher.Block, iv []byte, data []byte) []byte {
	dec := cipher.NewCFBDecrypter(block, iv)
	decrypted := make([]byte, len(data))
	dec.XORKeyStream(decrypted, data)
	return decrypted
}

func EncryptDecryptWithOFB(block cipher.Block, iv []byte, data []byte) []byte {
	enc := cipher.NewOFB(block, iv)
	encrypted := make([]byte, len(data))
	enc.XORKeyStream(encrypted, data)
	return encrypted
}

func EncryptDecryptWithCTR(block cipher.Block, iv []byte, data []byte) []byte {
	enc := cipher.NewCTR(block, iv)
	encrypted := make([]byte, len(data))
	enc.XORKeyStream(encrypted, data)
	return encrypted
}
