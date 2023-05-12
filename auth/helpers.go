package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
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

func getStringSha1(str string) string {
	hashBytes := sha1.Sum([]byte(str))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}
