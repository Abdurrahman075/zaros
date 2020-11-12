package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat file env")
	}

	result := os.Getenv(key)
	return result
}
