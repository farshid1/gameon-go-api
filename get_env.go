package gameon

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	projectDirName := "gameon-go-api"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}
