package worker

import (
	"context"
	"encoding/json"
	"final_project/internal/domain/appointment"
	leaverequests "final_project/internal/domain/leave_requests"
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
	settingRepo  setting.Repository
	redisRepo    redis.Repository
	leaveReqRepo leaverequests.Repository
}

func NewAppointmentCronJob(settingRepo setting.Repository, redisRepo redis.Repository, leaveReqRepo leaverequests.Repository) *AppointmentCronJob {
	return &AppointmentCronJob{
		settingRepo:  settingRepo,
		redisRepo:    redisRepo,
		leaveReqRepo: leaveReqRepo,
	}
}

func (c *AppointmentCronJob) ScheduleAppointment(ctx context.Context) error {
	settingValues := make(map[string]string, 0)

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

		settingValues[value] = tempSetting.Value
	}

	// Xử lí kiểm tra ngày hợp lệ
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	today := time.Now().In(loc)

	for i := 1; i < 356; i++ {
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

			isInLeaveTime, err := c.leaveReqRepo.IsInLeaveRequest(ctx, tomorrow)
			if err != nil {
				return err
			}

			if !isInLeaveTime {
				if err := c.createAppointment(ctx, tomorrow, settingValues); err != nil {
					return err
				}
			}
		} else if cal.IsWeekend(tomorrow) {
			fmt.Println(tomorrow.String() + " là ngày cuối tuần")
		} else if actual || observed {
			fmt.Println(tomorrow.String() + " là ngày lễ")
		} else {
			fmt.Println("📅 Không rõ trạng thái " + tomorrow.String())
		}
	}

	return nil
}

func (c *AppointmentCronJob) createAppointment(ctx context.Context, appointmentDay time.Time, settings map[string]string) error {
	var (
		itemClaimReqs           map[string]string
		appointmentPerMorning   int = 0
		appointmentPerAfternoon int = 0
		// startMorningTime        time.Time
		// endMorningTime          time.Time
		// startAfternoonTime      time.Time
		// endAfternoonTime        time.Time
	)

	appointmentPerMorning, _ = strconv.Atoi(settings["appointmentPerMorning"])
	appointmentPerAfternoon, _ = strconv.Atoi(settings["appointmentPerAfternoon"])

	startMorningTime, err := helpers.ParseToTime(settings["startMorningTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả:", startMorningTime.String())
	}

	endMorningTime, err := helpers.ParseToTime(settings["endMorningTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả:", endMorningTime.String())
	}

	startAfternoonTime, err := helpers.ParseToTime(settings["startAfternoonTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả:", startAfternoonTime.String())
	}

	endAfternoonTime, err := helpers.ParseToTime(settings["endAfternoonTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả:", endAfternoonTime.String())
	}

	newItemClaimReqs := make(map[string]warehouse.ClaimRequestItem, 0)
	userAppointment := make(map[uint]appointment.Appointment, 0)
	userAppointmentItem := make(map[uint][]appointment.AppointmentItem, 0)

	itemClaimReqs, err = c.redisRepo.GetAllFromRedisHash(ctx, enums.ItemClaimRequest)
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
			var (
				startTime time.Time
				endTime   time.Time
			)

			if appointmentPerMorning == 0 && appointmentPerAfternoon == 0 {
				break
			} else if appointmentPerMorning != 0 {
				startTime = time.Date(
					appointmentDay.Year(), appointmentDay.Month(), appointmentDay.Day(), // lấy ngày từ biến `day`
					startMorningTime.Hour(), startMorningTime.Minute(), startMorningTime.Second(), 0, // lấy giờ từ biến `hourOnly`
					appointmentDay.Location(), // giữ nguyên location theo `day`
				)

				endTime = time.Date(
					appointmentDay.Year(), appointmentDay.Month(), appointmentDay.Day(), // lấy ngày từ biến `day`
					endMorningTime.Hour(), endMorningTime.Minute(), endMorningTime.Second(), 0, // lấy giờ từ biến `hourOnly`
					appointmentDay.Location(), // giữ nguyên location theo `day`
				)
			} else {
				startTime = time.Date(
					appointmentDay.Year(), appointmentDay.Month(), appointmentDay.Day(), // lấy ngày từ biến `day`
					startAfternoonTime.Hour(), startAfternoonTime.Minute(), startAfternoonTime.Second(), 0, // lấy giờ từ biến `hourOnly`
					appointmentDay.Location(), // giữ nguyên location theo `day`
				)

				endTime = time.Date(
					appointmentDay.Year(), appointmentDay.Month(), appointmentDay.Day(), // lấy ngày từ biến `day`
					endAfternoonTime.Hour(), endAfternoonTime.Minute(), endAfternoonTime.Second(), 0, // lấy giờ từ biến `hourOnly`
					appointmentDay.Location(), // giữ nguyên location theo `day`
				)
			}

			if itemClaimReq.ItemQuantity >= user.Quantity {
				userAppointment[user.ID] = appointment.Appointment{
					UserID:    user.ID,
					StartTime: startTime,
					EndTime:   endTime,
					Status:    1,
				}

				userAppointmentItem[user.ID] = append(userAppointmentItem[user.ID], appointment.AppointmentItem{
					ItemID:          uint(itemID), //item id
					ActualQuantity:  int(user.Quantity),
					MissingQuantity: 0,
				})

				itemClaimReq.Users = append(itemClaimReq.Users[:idx], itemClaimReq.Users[idx+1:]...)

				itemClaimReq.ItemQuantity -= user.Quantity
			} else {
				userAppointment[user.ID] = appointment.Appointment{
					UserID:    user.ID,
					StartTime: startTime,
					EndTime:   endTime,
					Status:    1,
				}

				userAppointmentItem[user.ID] = append(userAppointmentItem[user.ID], appointment.AppointmentItem{
					ItemID:          uint(itemID), //item id
					ActualQuantity:  int(itemClaimReq.ItemQuantity),
					MissingQuantity: int(user.Quantity - itemClaimReq.ItemQuantity),
				})

				itemClaimReq.Users[idx].Quantity = 0

				break
			}

			if appointmentPerMorning == 0 {
				appointmentPerAfternoon--
			} else {
				appointmentPerMorning--
			}
		}

		newItemClaimReqs[key] = itemClaimReq
	}

	// for key, value := range newItemClaimReqs {
	// 	newItemClaimReqJSON, err := json.Marshal(value)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	//Lưu danh sách người dùng đăng kí mới vào redis
	// 	if err := c.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, key, string(newItemClaimReqJSON)); err != nil {
	// 		return err
	// 	}
	// }

	//Lọc qua danh sách các user đã được thông qua đăng kí nhận đồ và cập nhật số lượng đồ
	// for key, value := range userAppointmentItem {
	// 	//Lấy ra danh sách các item của user đã đăng kí dưới dạng JSON
	// 	userClaimReqJSON, err := c.redisRepo.GetFromRedisHash(ctx, enums.UserClaimRequest, "user:"+strconv.Itoa(int(key)))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if userClaimReqJSON == "" {
	// 		break
	// 	}

	// 	var (
	// 		userClaimReqs []warehouse.CreateClaimRequestItem
	// 	)

	// 	userClaimReqMap := make(map[uint]uint, 0)

	// 	//Decode JSON thành mảng các món đồ đã đăng kí của user
	// 	err = json.Unmarshal([]byte(userClaimReqJSON), &userClaimReqs)
	// 	if err != nil {
	// 		return errors.New("Có lỗi khi thực hiện decode JSON: " + err.Error())
	// 	}

	// 	//Lặp qua danh sách các món đồ user đã đăng kí --> Map thành danh sách số lượng các sản phẩm user đã xin, key là id của item
	// 	for _, user := range userClaimReqs {
	// 		userClaimReqMap[user.ItemID] = user.Quantity
	// 	}

	// 	//Lặp qua danh sách các item_appointment và cập nhật lại số lượng trong map số lượng hiện tại
	// 	for _, item := range value {
	// 		if item.MissingQuantity == 0 {
	// 			delete(userClaimReqMap, item.ItemID)
	// 		} else {
	// 			userClaimReqMap[item.ItemID] = uint(item.MissingQuantity)
	// 		}
	// 	}

	// 	currentUserClaimReqs := make([]warehouse.CreateClaimRequestItem, 0)

	// 	//Lặp qua map và lưu lại danh sách các món đồ user chưa được nhận
	// 	for idx, quantity := range userClaimReqMap {
	// 		currentUserClaimReqs = append(currentUserClaimReqs, warehouse.CreateClaimRequestItem{
	// 			ItemID:   idx,
	// 			Quantity: quantity,
	// 		})
	// 	}

	// 	//Encode mảng các món đồ của user đăng kí hiện tại(sau khi tính toán)
	// 	currentUserClaimReqsJSON, err := json.Marshal(currentUserClaimReqs)
	// 	if err != nil {
	// 		return errors.New("Có lỗi khi thực hiện encode JSON: " + err.Error())
	// 	}

	// 	//Lưu vào redis
	// 	if err := c.redisRepo.SetToRedisHash(ctx, enums.UserClaimRequest, "user:"+strconv.Itoa(int(key)), string(currentUserClaimReqsJSON)); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
