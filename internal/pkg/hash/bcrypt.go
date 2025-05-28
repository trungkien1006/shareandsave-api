package hash

import (
	"crypto/sha256"
	"encoding/base32"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// HashEmailPhone tạo chuỗi hash 8 ký tự từ email + sdt
func HashEmailPhone(email, phone string) string {
	data := email + phone
	hash := sha256.Sum256([]byte(data))
	// Dùng base32 để ra ký tự chữ + số, loại bỏ padding '='
	encoded := strings.TrimRight(base32.StdEncoding.EncodeToString(hash[:]), "=")
	// Lấy 8 ký tự đầu
	if len(encoded) > 8 {
		return encoded[:8]
	}
	return encoded
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // DefaultCost = 10

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
