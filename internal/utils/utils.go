package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func GenerateSlug(checkExists func(string) bool, length, attempts int) (string, error) {
	for i := 0; i < attempts; i++ {
		sb := strings.Builder{}
		for j := 0; j < length; j++ {
			sb.WriteByte(alphabet[rand.Intn(len(alphabet))])
		}
		slug := sb.String()
		if !checkExists(slug) {
			return slug, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique slug after %d tries", attempts)
}
