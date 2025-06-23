package worker

import (
	"context"
	"encoding/json"
	"errors"
	"final_project/internal/domain/appointment"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/setting"
	"final_project/internal/domain/warehouse"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rickar/cal/v2"
)

type AppointmentCronJob struct {
	settingRepo setting.Repository
	redisRepo   redis.Repository
}

func NewAppointmentCronJob(settingRepo setting.Repository, redisRepo redis.Repository) *AppointmentCronJob {
	return &AppointmentCronJob{
		settingRepo: settingRepo,
		redisRepo:   redisRepo,
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

	// Xử lí kiểm tra ngày hợp lệ
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	today := time.Now().In(loc)

	for i := 1; i < 15; i++ {
		tomorrow := today.AddDate(0, 0, 1)

		// Tạo lịch Việt Nam cho năm hiện tại
		year := tomorrow.Year()
		calendar := helpers.VietnamHolidayCalendar(year)

		// Thêm tên mô tả nếu cần
		calendar.Name = "Lịch Việt Nam"
		calendar.Description = "Lịch làm việc và nghỉ lễ theo Bộ luật Lao động"

		// Kiểm tra ngày hôm nay
		actual, observed, _ := calendar.IsHoliday(tomorrow)

		if calendar.IsWorkday(tomorrow) {
			fmt.Println(tomorrow.String() + " là ngày làm việc")

			break
		} else if cal.IsWeekend(tomorrow) {
			fmt.Println(tomorrow.String() + " là ngày cuối tuần")
		} else if actual || observed {
			fmt.Println(tomorrow.String() + " là ngày lễ")
		} else {
			fmt.Println("📅 Không rõ trạng thái " + tomorrow.String())
		}
	}

	if err := c.createAppointment(ctx); err != nil {
		return err
	}

	return nil
}

func (c *AppointmentCronJob) createAppointment(ctx context.Context) error {
	var itemClaimReqs map[string]string

	userClaimResult := make(map[uint][]appointment.AppointmentItem, 0)

	itemClaimReqs, err := c.redisRepo.GetAllFromRedisHash(ctx, enums.ItemClaimRequest)
	if err != nil {
		return err
	}

	//Lọc qua danh sách các item
	for key, value := range itemClaimReqs {
		var itemClaimReq warehouse.ClaimRequestItem

		err := json.Unmarshal([]byte(value), &itemClaimReq)
		if err != nil {
			return err
		}

		itemIDStr := strings.Split(key, ":")[1]

		itemID, _ := strconv.Atoi(itemIDStr)

		//Lọc qua từng user đã đăng kí nhận đồ trong item
		for idx, user := range itemClaimReq.Users {
			if itemClaimReq.ItemQuantity >= user.Quantity {
				userClaimResult[user.ID] = append(userClaimResult[user.ID], appointment.AppointmentItem{
					ItemID:          uint(itemID), //item id
					ActualQuantity:  int(user.Quantity),
					MissingQuantity: 0,
				})

				itemClaimReq.Users = append(itemClaimReq.Users[:idx], itemClaimReq.Users[idx+1:]...)

				itemClaimReq.ItemQuantity -= user.Quantity
			} else {
				userClaimResult[user.ID] = append(userClaimResult[user.ID], appointment.AppointmentItem{
					ItemID:          uint(itemID), //item id
					ActualQuantity:  int(itemClaimReq.ItemQuantity),
					MissingQuantity: int(user.Quantity - itemClaimReq.ItemQuantity),
				})

				itemClaimReq.Users[idx].Quantity = 0

				break
			}
		}

		newItemClaimReqJSON, err := json.Marshal(itemClaimReq)
		if err != nil {
			return err
		}

		//Lưu danh sách người dùng đăng kí mới vào redis
		if err := c.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, "item:"+key, string(newItemClaimReqJSON)); err != nil {
			return err
		}
	}

	//Lọc qua danh sách các user đã được thông qua đăng kí nhận đồ và cập nhật số lượng đồ
	for key, value := range userClaimResult {
		//Lấy ra danh sách các item của user đã đăng kí dưới dạng JSON
		userClaimReqJSON, err := c.redisRepo.GetFromRedisHash(ctx, enums.UserClaimRequest, "user:"+strconv.Itoa(int(key)))
		if err != nil {
			return err
		}

		if userClaimReqJSON == "" {
			break
		}

		var (
			userClaimReqs []warehouse.CreateClaimRequestItem
		)

		userClaimReqMap := make(map[uint]uint, 0)

		//Decode JSON thành mảng các món đồ đã đăng kí của user
		err = json.Unmarshal([]byte(userClaimReqJSON), &userClaimReqs)
		if err != nil {
			return errors.New("Có lỗi khi thực hiện decode JSON: " + err.Error())
		}

		//Lặp qua danh sách các món đồ user đã đăng kí --> Map thành danh sách số lượng các sản phẩm user đã xin, key là id của item
		for _, user := range userClaimReqs {
			userClaimReqMap[user.ItemID] = user.Quantity
		}

		//Lặp qua danh sách các item_appointment và cập nhật lại số lượng trong map số lượng hiện tại
		for _, item := range value {
			if item.MissingQuantity == 0 {
				delete(userClaimReqMap, item.ItemID)
			} else {
				userClaimReqMap[item.ItemID] = uint(item.MissingQuantity)
			}
		}

		currentUserClaimReqs := make([]warehouse.CreateClaimRequestItem, 0)

		//Lặp qua map và lưu lại danh sách các món đồ user chưa được nhận
		for idx, quantity := range userClaimReqMap {
			currentUserClaimReqs = append(currentUserClaimReqs, warehouse.CreateClaimRequestItem{
				ItemID:   idx,
				Quantity: quantity,
			})
		}

		//Encode mảng các món đồ của user đăng kí hiện tại(sau khi tính toán)
		currentUserClaimReqsJSON, err := json.Marshal(currentUserClaimReqs)
		if err != nil {
			return errors.New("Có lỗi khi thực hiện encode JSON: " + err.Error())
		}

		//Lưu vào redis
		if err := c.redisRepo.SetToRedisHash(ctx, enums.UserClaimRequest, "user:"+strconv.Itoa(int(key)), string(currentUserClaimReqsJSON)); err != nil {
			return err
		}
	}

	return nil
}
