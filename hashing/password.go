package hashing

import (
	"crypto/sha1"
	"fmt"
)

// Хешер пароля с использованием заданной соли
type SHA1Hasher struct {
	salt string
}

// NewSHA1Hasher возвращает новый хешер
func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

// Hash хеширует пароль с использованием соли
func (h SHA1Hasher) Hash(password string) (string, error) {
	hash := sha1.New()

	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}
