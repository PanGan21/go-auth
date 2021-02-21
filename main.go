package main

import (
	"context"
	"go-auth/auth"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file detected.")
	}
}

// NewRedisDB return a new redis client
func NewRedisDB(host string, port string, password string) (*redis.Client) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
		Password: password,
		DB: 0,
	})
	return redisClient
}

func main() {
	appAddr := ":" + os.Getenv("PORT")

	// setup redis
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient := NewRedisDB(redisHost, redisPort, redisPassword)

	var rd = auth.NewAuth(redisClient)
	var tk = auth.NewToken()

	var router = gin.Default()

	srv := &http.Server{
		Addr: appAddr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen: %s\n", err)
		}
	}()
	//	Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting.")
}