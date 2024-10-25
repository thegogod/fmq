package env

import "os"

func GetOrDefault(name string, defaultValue string) string {
	v := os.Getenv(name)

	if v == "" {
		return defaultValue
	}

	return v
}
