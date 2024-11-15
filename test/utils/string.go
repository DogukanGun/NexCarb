package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	// Define the character set
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate the random string
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
