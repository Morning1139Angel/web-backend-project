package main

import (
	"context"
	"net/http"
	"os"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

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

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func addIpToBlockList(rds *redis.Client) func(*gin.Context, ratelimit.Info) {
	return func(c *gin.Context, info ratelimit.Info) {
		// Add the IP to Redis with a blocking duration of 24 hours
		err := rds.Set(context.Background(), c.ClientIP(), true, 24*time.Hour).Err()
		if err != nil {
			// Handle Redis error
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}

func RequestCountMiddleware(rds *redis.Client, maxRequestPerSec uint) gin.HandlerFunc {
	store := ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: rds,
		Rate:        time.Second,
		Limit:       maxRequestPerSec,
	})

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: addIpToBlockList(rds),
		KeyFunc:      keyFunc,
	})

	return mw
}

func IPBlockMiddleware(rds *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		// Check if the IP is blocked
		_, err := rds.Get(context.Background(), ip).Result()
		if err != nil && err != redis.Nil {
			// Handle Redis error
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		} else if err != nil && err == redis.Nil {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Your IP is blocked. Access denied.",
			})
		}
	}
}
