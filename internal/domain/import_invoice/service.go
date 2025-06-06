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

func (s *ImportInvoiceService) GenerateSKU(itemID int) string {
	nowDate := GetCurrentTimeVN().Format("2006-01-02")
	// Tạo số ngẫu nhiên 4 chữ số (0000-9999)
	rand.Seed(time.Now().UnixNano())  // Đảm bảo ngẫu nhiên mỗi lần gọi
	randomCounter := rand.Intn(10000) // Số ngẫu nhiên từ 0 đến 9999
	return fmt.Sprintf("%d-%s-%04d", itemID, nowDate, randomCounter)
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
