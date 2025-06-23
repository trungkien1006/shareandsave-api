package worker

import (
	"context"
	"final_project/internal/domain/setting"
)

type AppointmentCronJob struct {
	settingRepo setting.Repository
}

func NewAppointmentCronJob(settingRepo setting.Repository) *AppointmentCronJob {
	return &AppointmentCronJob{
		settingRepo: settingRepo,
	}
}

func (c *AppointmentCronJob) ScheduleAppointment(ctx context.Context) error {
	var (
	// appointmentPerDay int = 0
	// startTime         time.Time
	// endTime           time.Time
	)

	settingValues := make(map[string]setting.Setting, 0)

	settingKeys := []string{
		"appointmentPerMorning",
		"appointmentPerAfternoon",
		"startMorningTime",
		"endMorningTime",
		"startAfternoonTime",
		"endAfternoonTime",
	}

	for _, value := range settingKeys {
		var tempSetting setting.Setting

		err := c.settingRepo.GetByKey(ctx, &tempSetting, value)
		if err != nil {
			return err
		}

		settingValues[value] = tempSetting
	}

	// appointmentPerDay, err = strconv.Atoi(appointmentPerDaySetting.Value)
	// if err != nil {
	// 	return errors.New("Có lỗi khi chuyển đổi giá trị tổng số lich hẹn 1 ngày: " + err.Error())
	// }

	// startTime, err = helpers.ParseToTime(startTimeSetting.Value)
	// if err != nil {
	// 	fmt.Println("Lỗi parse thời gian:", err)
	// } else {
	// 	fmt.Println("Kết quả:", startTime.String())
	// }

	// endTime, err = helpers.ParseToTime(endTimeSetting.Value)
	// if err != nil {
	// 	fmt.Println("Lỗi parse thời gian:", err)
	// } else {
	// 	fmt.Println("Kết quả:", endTime.String())
	// }

	//Xử lí kiểm tra ngày hợp lệ
	// loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	// t := time.Now().In(loc)

	// Tạo BusinessCalendar
	// c := cal.NewBusinessCalendar()
	// c.Name = "MyCompany Calendar"
	// c.Description = "Lịch làm việc và nghỉ lễ"

	// // Thêm ngày lễ (ở đây lấy ví dụ Mỹ, bạn có thể tự định nghĩa ngày lễ VN)
	// c.AddHoliday(
	// 	us.NewYear,
	// 	us.MemorialDay,
	// 	us.IndependenceDay,
	// 	us.LaborDay,
	// 	us.ThanksgivingDay,
	// 	us.ChristmasDay,
	// )

	// // Thiết lập tuần làm việc (thứ 2–6)
	// c.SetWorkday(time.Saturday, false)
	// c.SetWorkday(time.Sunday, false)

	// // Kiểm tra ngày hôm nay có phải là ngày làm việc?
	// if c.IsWorkday(t) {
	// 	fmt.Println("✅ Hôm nay là ngày làm việc")
	// } else if cal.IsWeekend(t) {
	// 	fmt.Println("⛱ Cuối tuần")
	// } else if c.IsHoliday(t) {
	// 	fmt.Println("🎉 Hôm nay là ngày lễ")
	// }

	return nil
}
