package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Conf(key string)string  {
	if err:=godotenv.Load(".env");err!=nil{
		log.Print(err)
	}
	return os.Getenv(key)
}
