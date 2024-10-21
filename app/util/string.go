package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func MakeRandomString(length int) string {
	randBytes := make([]byte, length)
	for {
		if _, err := io.ReadFull(rand.Reader, randBytes); err == nil {
			break
		}
	}
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}

func Hash(s string, count int) string {
	hash := sha256.Sum256([]byte(s))
	for i := 0; i < count; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return fmt.Sprintf("%x", hash)
}
