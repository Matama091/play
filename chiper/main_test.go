package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"testing"
)

func call(b *testing.B, dataSize int, f func(block cipher.Block, iv []byte, data []byte) []byte) ([]byte, error) {
	b.StopTimer()
	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, 16) // = aes.BlockSize()
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	data := make([]byte, dataSize)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return nil, err
	}
	b.StartTimer()
	return f(block, iv, data), nil
}

func BenchmarkShortEncryptWithCBC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 128, EncryptWithCBC)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkShortDecryptWithCBC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 128, DecryptWithCBC)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkShortEncryptWithCFB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 128, EncryptWithCFB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkShortDecryptWithCFB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 128, DecryptWithCFB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkShortEncryptDecryptWithOFB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 128, EncryptDecryptWithOFB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkShortEncryptDecryptWithCTR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 128, EncryptDecryptWithCTR)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLongEncryptWithCBC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 12800, EncryptWithCBC)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLongDecryptWithCBC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 12800, DecryptWithCBC)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLongEncryptWithCFB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 12800, EncryptWithCFB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLongDecryptWithCFB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 12800, DecryptWithCFB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLongEncryptDecryptWithOFB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 12800, EncryptDecryptWithOFB)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLongEncryptDecryptWithCTR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := call(b, 12800, EncryptDecryptWithCTR)
		if err != nil {
			b.Error(err)
		}
	}
}
