package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT       = 0
	DBURL      = "mongodb+srv://sanket:Sanket123@service.7rseb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	REDIS_HOST = ""
	DB_INDEX   = 0
	EXP        = 0
)

func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error is : ", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 8080
	}
	REDIS_HOST = os.Getenv("REDIS_HOST")
	DB_INDEX, err = strconv.Atoi(os.Getenv("DB_INDEX"))
	if err != nil {
		DB_INDEX = 1
	}
	EXP, err = strconv.Atoi(os.Getenv("EXP"))
	if err != nil {
		EXP = 600
	}

}
