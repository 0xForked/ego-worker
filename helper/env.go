package helper

import (
	"github.com/joho/godotenv"
	"os"
)

// Get env variable from file using godotenv library
func GetEnv(key, defaultValue string) string {
	// read .env file with godotenv
	err := godotenv.Load()
	// validate error
	CheckError(err, "Error loading .env file")
	// retrieves the value of the environment variable named by the key.
	value := os.Getenv(key)
	// validate env value
	if len(value) == 0 {
		// if value size <= 0 return default value
		ShowMessage(defaultValue)
		return defaultValue
	}
	// return value if key value size > 0
	return value
}
