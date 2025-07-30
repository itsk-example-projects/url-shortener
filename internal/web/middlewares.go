package web

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"
)

func BasicAuth(r *http.Request) (string, string) {
	auth := r.Header.Get("Authorization")
	parts := strings.SplitN(auth, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		return "", ""
	}
	b, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", ""
	}
	creds := strings.SplitN(string(b), ":", 2)
	if len(creds) != 2 {
		return "", ""
	}
	return creds[0], creds[1]
}

func SetUserIDCookie(w http.ResponseWriter, uid string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "uid",
		Value:    uid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   4 * 365 * 24 * 60 * 60, // 4 years
	})
}

func GetUserID(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("uid")
	if err == nil && cookie.Value != "" {
		return cookie.Value
	}
	uid, err := randomString(16)
	if err != nil {
		log.Printf("failed to generate random string for user ID: %v", err)
		uid = fmt.Sprintf("fallback-%d", time.Now().UnixNano()) //TODO: _
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "uid",
		Value:    uid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   4 * 365 * 24 * 60 * 60, // 4 years
	})
	return uid
}

func randomString(length int) (string, error) {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		result[i] = chars[num.Int64()]
	}
	return string(result), nil
}
