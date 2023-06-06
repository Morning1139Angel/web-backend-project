package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"os"

	"github.com/go-redis/redis/v8"
)

func GenerateNonce(len int) string {
	nonceByte := make([]byte, 16)
	if _, err := rand.Read(nonceByte); err != nil {
		panic(err)
	}

	// Encode the byte slice as a base64-encoded string of length 20
	nonceStr := base64.RawURLEncoding.EncodeToString(nonceByte)[:20]
	return nonceStr
}

func GetStringSha1(str string) string {
	hashBytes := sha1.Sum([]byte(str))
	hashString := hex.EncodeToString(hashBytes[:])
	return hashString
}

func InitRedicClient() *redis.Client {
	var redisHost = os.Getenv("REDIS_HOST")
	var redisPort = os.Getenv("REDIS_PORT")
	var redisPassword = os.Getenv("REDIS_PASSWORD")
	return redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}

func GenerateRandomOddNumber() (uint64, error) {
	for {
		// Generate a random 64-bit number
		randomNum, err := rand.Int(rand.Reader, new(big.Int).SetBit(new(big.Int), 64, 1))
		if err != nil {
			return 0, err
		}

		// Check if the number is odd
		if randomNum.Bit(0) == 1 {
			return randomNum.Uint64(), nil
		}
	}
}

func StorageKeyFromNonces(clientNonce string, nonceServer string) string {
	key := GetStringSha1(clientNonce + nonceServer)
	return key
}
