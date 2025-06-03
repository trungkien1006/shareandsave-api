package importinvoice

import (
	"fmt"
	"math/rand"
	"time"
)

type ImportInvoiceService struct{}

func NewImportInvoiceService() *ImportInvoiceService {
	return &ImportInvoiceService{}
}

func (s *ImportInvoiceService) GenerateSKU(itemID int, counter int64) string {
	nowDate := GetCurrentTimeVN().Format("2006-01-02")

	return string(itemID) + "-" + nowDate + "-" + string(counter)
}

func GetCurrentTimeVN() time.Time {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("⚠️ Lỗi khi load location:", err)
		return time.Now().UTC() // Fallback về UTC nếu load location lỗi
	}

	return time.Now().In(location)
}

func (s *ImportInvoiceService) GenerateUniqueDigitString(length int) (string, error) {
	if length > 10 || length < 1 {
		return "", fmt.Errorf("độ dài phải từ 1 đến 10")
	}

	digits := []rune("0123456789")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(digits), func(i, j int) {
		digits[i], digits[j] = digits[j], digits[i]
	})

	return string(digits[:length]), nil
}
