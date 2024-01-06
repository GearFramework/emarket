package gear

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

const (
	bcryptDefaultCost = 13
	letterBytes       = "abcdefghijklmnopqrstuvwxyz@#+1234567890*&%ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits     = 6
	letterIdxMask     = 63
	letterIdxMax      = 10
)

// GenerateRandomPassword генерация случайной строки указанной длинны
func GenerateRandomPassword(lenPass int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(lenPass)
	for i, cache, remain := lenPass-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return sb.String()
}

// HashPassword создание хэша для указанного пароля
func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcryptDefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ValidatePassword сравнение переданного пароля с хэшем
func ValidatePassword(hash, pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		return false
	}
	return true
}
