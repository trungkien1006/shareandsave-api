package helpers

import (
	"fmt"
	"strings"
	"time"
)

func FormatDurationVi(d time.Duration) string {
	if d < 0 {
		d = 0
	}

	totalMinutes := int(d.Minutes())
	days := totalMinutes / (60 * 24)
	hours := (totalMinutes % (60 * 24)) / 60
	minutes := totalMinutes % 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d ngày", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d giờ", hours))
	}
	if minutes > 0 || len(parts) == 0 {
		parts = append(parts, fmt.Sprintf("%d phút", minutes))
	}

	return strings.Join(parts, " ")
}

func GetCurrentTimeVN() time.Time {
	location, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("⚠️ Lỗi khi load location:", err)
		return time.Now().UTC() // Fallback về UTC nếu load location lỗi
	}

	return time.Now().In(location)
}

func IsPast(dateStr string) (bool, error) {
	// Định dạng thời gian theo yêu cầu: yyyy-MM-ddThh:mm
	const layout = "2006-01-02T15:04"

	// Chuyển đổi chuỗi thành kiểu time.Time
	inputTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return false, err
	}

	// Lấy thời gian hiện tại
	now := time.Now()

	// So sánh với thời gian hiện tại
	return inputTime.Before(now), nil
}

func ParseToTime(input string) (time.Time, error) {
	layout := "15:04"
	return time.Parse(layout, input)
}
