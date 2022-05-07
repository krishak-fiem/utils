package env

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load the env file: %v", err)
	}
}

func Get(key string) string {
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()
	return myEnv[key]
}
