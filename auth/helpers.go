package main

import (
	"crypto/rand"
	"encoding/base64"
)

func generateNonce(len int) string {
	nonceByte := make([]byte, 16)
	if _, err := rand.Read(nonceByte); err != nil {
		panic(err)
	}

	// Encode the byte slice as a base64-encoded string of length 20
	nonceStr := base64.RawURLEncoding.EncodeToString(nonceByte)[:20]
	return nonceStr
}
