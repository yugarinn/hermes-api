package utils

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "hermes-api"

func LoadEnvFile(env string) {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
    currentWorkDirectory, _ := os.Getwd()
    rootPath := projectName.Find([]byte(currentWorkDirectory))

	var err error

	if env == "production" {
		err = godotenv.Load(string(rootPath) + `/.env`)
	}

	if env == "development" {
		err = godotenv.Load(string(rootPath) + `/.env.development`)
	}

	if env == "test" {
		err = godotenv.Load(string(rootPath) + `/.env.test`)
	}

    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
}

func IsProduction() bool {
	return os.Getenv("HERMES_ENV") == "production"
}

func IsDevelopment() bool {
	return os.Getenv("HERMES_ENV") == "development"
}

func IsTest() bool {
	return os.Getenv("HERMES_ENV") == "test"
}
