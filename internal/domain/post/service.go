package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

// GenerateSlug tạo slug từ title
func (s *PostService) GenerateSlug(title string) string {
	// Chuyển title về chữ thường, thay thế ký tự đặc biệt bằng dấu gạch ngang
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := strings.ToLower(title)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	// Sinh số ngẫu nhiên 3 chữ số
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(900) + 100 // 100-999

	return fmt.Sprintf("%s-%d", slug, id)
}

// GenerateContent tạo content từ info
func (s *PostService) GenerateContent(info string) (string, error) {
	input := []byte(info)

	// Unmarshal vào map
	var rawMap map[string]string
	if err := json.Unmarshal(input, &rawMap); err != nil {
		return "", errors.New("Lỗi khi mã hóa JSON: " + err.Error())
	}

	// Duyệt qua map và gom value
	var values []string
	for _, v := range rawMap {
		values = append(values, v)
	}

	// In ra kết quả
	result, _ := json.Marshal(values)

	return string(result), nil
}
