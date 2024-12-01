package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
