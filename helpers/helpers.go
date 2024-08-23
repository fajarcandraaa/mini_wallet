package helpers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"
	"unicode"

	"github.com/fajarcandraaa/mini_wallet/internal/entity"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is func to encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is func to validate password hasher
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsNumber is func to validate type from input string is number or alphabet
func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func GenerateHexadecimalStringTokent() (*string, error) {
	hashLength := 20 // 20 bytes = 40 hexadecimal characters

	// Generate random bytes
	randomBytes := make([]byte, hashLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Convert random bytes to hexadecimal string
	hashString := hex.EncodeToString(randomBytes)
	return &hashString, nil
}

func ParseTokenHex(token string) (string, error) {
	substrings := strings.Split(token, " ")
	trimmedStr := strings.TrimSpace(substrings[1])

	return trimmedStr, nil
}

func FindCustomerXidFromToken(ctx context.Context, rds *redis.Client, token string) (string, error) {
	var (
		tokenKey string
	)
	keys, err := rds.Keys(ctx, "*").Result()
	if err != nil {
		return "", err
	}

	for _, k := range keys {
		v, err := rds.Get(ctx, k).Result()
		if err != nil {
			return "", err
		}
		if v == token {
			tokenKey = k
			break
		}
	}

	if tokenKey == "" {
		return "", entity.ErrPermissionNotAllowed
	}

	return tokenKey, nil
}

func GetCustomerXidFromToken(rdsKey string) string {
	substrings := strings.Split(rdsKey, " ")
	trimmedStr := strings.TrimSpace(substrings[1])

	return trimmedStr
}
