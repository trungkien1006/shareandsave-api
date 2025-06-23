package worker

// import (
// 	"context"
// 	"errors"
// 	"final_project/internal/domain/setting"
// 	"final_project/internal/pkg/helpers"
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"google.golang.org/genproto/googleapis/rpc/context"
// )

// type AppointmentCronJob struct {
// 	settingRepo setting.Repository
// }

// func NewAppointmentCronJob(settingRepo setting.Repository) *AppointmentCronJob {
// 	return &AppointmentCronJob{
// 		settingRepo: settingRepo,
// 	}
// }

// func (c *AppointmentCronJob) ScheduleAppointment(ctx context.Context) error {
// 	var (
// 		appointmentPerDay        int = 0
// 		startTime                time.Time
// 		endTime                  time.Time
// 		appointmentPerDaySetting setting.Setting
// 		startTimeSetting         setting.Setting
// 		endTimeSetting           setting.Setting
// 	)

// 	err := c.settingRepo.GetByKey(ctx, &appointmentPerDaySetting, "appointmentPerDay")
// 	if err != nil {
// 		return err
// 	}

// 	err = c.settingRepo.GetByKey(ctx, &startTimeSetting, "startTime")
// 	if err != nil {
// 		return err
// 	}

// 	err = c.settingRepo.GetByKey(ctx, &endTimeSetting, "endTime")
// 	if err != nil {
// 		return err
// 	}

// 	appointmentPerDay, err = strconv.Atoi(appointmentPerDaySetting.Value)
// 	if err != nil {
// 		return errors.New("Có lỗi khi chuyển đổi giá trị tổng số lich hẹn 1 ngày: " + err.Error())
// 	}

// 	startTime, err = helpers.ParseToTime(startTimeSetting.Value)
// 	if err != nil {
// 		fmt.Println("Lỗi parse thời gian:", err)
// 	} else {
// 		fmt.Println("Kết quả:", startTime.String())
// 	}

// 	endTime, err = helpers.ParseToTime(endTimeSetting.Value)
// 	if err != nil {
// 		fmt.Println("Lỗi parse thời gian:", err)
// 	} else {
// 		fmt.Println("Kết quả:", endTime.String())
// 	}

// 	return nil
// }
