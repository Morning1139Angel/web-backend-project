package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/go-redis/redis/v8"
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

func initRedicClient() *redis.Client {
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPort = os.Getenv("REDIS_PORT")
	var redisPassword = os.Getenv("REDIS_PASSWORD")
	return redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}

func generatePandG() (uint64, uint64) {
	return 0, 0
}
