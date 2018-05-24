package envkit

import (
	"os"
	"strconv"
)

func GetString(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) > 0 {
		return value
	}

	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err == nil {
		return i
	}

	return defaultValue

}
