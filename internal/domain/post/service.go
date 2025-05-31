package post

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

// GenerateSlug tạo slug từ title
func removeDiacritics(str string) string {
	// Normalize Unicode về dạng NFD (phân tách dấu)
	t := norm.NFD.String(str)
	// Loại bỏ các ký tự không phải chữ cái cơ bản
	sb := strings.Builder{}
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue // Mn = Mark, nonspacing (dấu)
		}
		sb.WriteRune(r)
	}
	return sb.String()
}

func (s *PostService) GenerateSlug(title string) string {
	// B1: Bỏ dấu tiếng Việt
	title = removeDiacritics(title)

	// B2: Chuyển lowercase
	slug := strings.ToLower(title)

	// B3: Thay các ký tự không phải chữ cái/số bằng dấu gạch ngang
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug = re.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")

	// B4: Sinh số ngẫu nhiên 3 chữ số để tránh trùng
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(900) + 100 // từ 100 đến 999

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
