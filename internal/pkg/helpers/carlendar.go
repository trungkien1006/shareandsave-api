package helpers

import (
	"fmt"
	"time"

	"github.com/6tail/lunar-go/calendar"
	"github.com/rickar/cal/v2"
)

// Trả về calendar Việt Nam đầy đủ cho năm chỉ định
func VietnamHolidayCalendar(year int) *cal.BusinessCalendar {
	c := cal.NewBusinessCalendar()
	c.SetWorkday(time.Saturday, false)
	c.SetWorkday(time.Sunday, false)

	// Ngày lễ dương lịch cố định
	fixed := []struct {
		name  string
		month time.Month
		day   int
	}{
		{"Tết Dương lịch", time.January, 1},
		{"Ngày giải phóng miền Nam", time.April, 30},
		{"Quốc tế Lao động", time.May, 1},
		{"Quốc khánh", time.September, 2},
	}
	for _, f := range fixed {
		addWithSubstitute(c, year, f.name, f.month, f.day)
	}

	// Giỗ tổ Hùng Vương 10/3 âm lịch
	lunarHG := calendar.NewLunarFromYmd(year, 3, 10)
	sHG := lunarHG.GetSolar()
	addWithSubstitute(c, year, "Giỗ tổ Hùng Vương", time.Month(sHG.GetMonth()), sHG.GetDay())

	// Tết Nguyên Đán: mùng 1–3 âm lịch (tự động chuyển sang dương)
	solarStart := calendar.NewLunarFromYmd(year, 1, 1).GetSolar()
	base := time.Date(solarStart.GetYear(), time.Month(solarStart.GetMonth()), solarStart.GetDay(), 0, 0, 0, 0, time.Local)
	for i := 0; i < 3; i++ {
		d := base.AddDate(0, 0, i)
		addWithSubstitute(c, year, fmt.Sprintf("Tết Nguyên Đán (mùng %d)", i+1), d.Month(), d.Day())
	}

	return c
}

// Hàm helper thêm ngày lễ + ngày nghỉ bù
func addWithSubstitute(c *cal.BusinessCalendar, year int, name string, month time.Month, day int) {
	c.AddHoliday(&cal.Holiday{Name: name, Month: month, Day: day})

	date := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	switch date.Weekday() {
	case time.Saturday:
		d2 := date.AddDate(0, 0, 2)
		c.AddHoliday(&cal.Holiday{Name: name + " (nghỉ bù)", Month: d2.Month(), Day: d2.Day()})
	case time.Sunday:
		d2 := date.AddDate(0, 0, 1)
		c.AddHoliday(&cal.Holiday{Name: name + " (nghỉ bù)", Month: d2.Month(), Day: d2.Day()})
	}
}
