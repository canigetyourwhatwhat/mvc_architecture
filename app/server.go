package app

import (
	"github.com/joho/godotenv"
	"log"
	"mvc_go/app/controllers"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = controllers.Server{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	dbConfig.DBUsername = getEnv("DB_USER", "test")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "test")
	dbConfig.DBName = getEnv("DB_NAME", "mvc_approach")

	server.InitServer(dbConfig)
	server.Run(":" + "8888")
}
