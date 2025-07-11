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
	settingRepo     setting.Repository
	redisRepo       redis.Repository
	appointmentRepo appointment.Repository
}

func NewAppointmentCronJob(settingRepo setting.Repository, redisRepo redis.Repository, appointmentRepo appointment.Repository) *AppointmentCronJob {
	return &AppointmentCronJob{
		settingRepo:     settingRepo,
		redisRepo:       redisRepo,
		appointmentRepo: appointmentRepo,
	}
}

func (c *AppointmentCronJob) ScheduleAppointment(ctx context.Context) error {
	settingValues := make(map[string]string, 0)

	settingKeys := []string{
		"numPerHour",
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
	today := time.Now()

	if err := c.createAppointment(ctx, today, settingValues); err != nil {
		return err
	}

	return nil
}

func (c *AppointmentCronJob) checkValidDay(ctx context.Context, today time.Time, start, end int, numPerHourCan int, endAfternoonTime, startAfternoonTime, endMorningTime time.Time) (time.Time, int, int, int, error) {
	loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	today = today.In(loc)

	for i := 1; i < 356; i++ {
		tomorrow := today.AddDate(0, 0, i)

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

			tempStart := start
			tempEnd := end

			for j := 0; j < 12; j++ {
				if tempEnd+j > endMorningTime.Hour() {
					tempStart = startAfternoonTime.Hour()
					tempEnd = tempStart + 1
				} else {
					tempStart = start + j
					tempEnd = end + j
				}

				startTime := time.Date(
					tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
					tempStart, 0, 0, 0,
					tomorrow.Location(), // hoặc timePart.Location() nếu cần
				)
				endTime := time.Date(
					tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
					tempEnd, 0, 0, 0,
					tomorrow.Location(), // hoặc timePart.Location() nếu cần
				)

				if endTime.Hour() > endAfternoonTime.Hour() {
					break
				}

				numPerHour, _ := c.appointmentRepo.IsInDay(ctx, startTime, endTime)

				if numPerHour < numPerHourCan {
					fmt.Println(strconv.Itoa(tempStart) + " là giờ bắt đầu")
					fmt.Println(strconv.Itoa(tempEnd) + " là giờ kết thúc")

					return tomorrow, numPerHourCan - numPerHour, startTime.Hour(), endTime.Hour(), nil
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

	return today, 0, 0, 0, errors.New("No valid day")
}

func (c *AppointmentCronJob) createAppointment(ctx context.Context, appointmentDay time.Time, settings map[string]string) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic in createAppointment: %v", r)
		}
	}()

	var (
		itemClaimReqs map[string]string
		numPerHour    int = 0
		err           error
	)

	numPerHour, err = strconv.Atoi(settings["numPerHour"])
	if err != nil {
		return err
	}

	startMorningTime, err := helpers.ParseToTime(settings["startMorningTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả start morning time:", startMorningTime.String())
	}

	endMorningTime, err := helpers.ParseToTime(settings["endMorningTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả end morning time:", endMorningTime.String())
	}

	startAfternoonTime, err := helpers.ParseToTime(settings["startAfternoonTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả start afternoon time:", startAfternoonTime.String())
	}

	endAfternoonTime, err := helpers.ParseToTime(settings["endAfternoonTime"])
	if err != nil {
		fmt.Println("Lỗi parse thời gian:", err)

		return err
	} else {
		fmt.Println("Kết quả end afternoon time:", endAfternoonTime.String())
	}

	tempNumPerHour := numPerHour
	currentStartHour := startMorningTime.Hour()
	currentEndHour := currentStartHour + 1

	appointmentDay, tempNumPerHour, currentStartHour, currentEndHour, err = c.checkValidDay(ctx, appointmentDay, startMorningTime.Hour(), startMorningTime.Hour()+1, tempNumPerHour, endAfternoonTime, startAfternoonTime, endMorningTime)
	if err != nil {
		return err
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
		var (
			itemClaimReq warehouse.ClaimRequestItem
		)

		err := json.Unmarshal([]byte(value), &itemClaimReq)
		if err != nil {
			return err
		}

		itemIDStr := strings.Split(key, ":")[1]

		itemID, _ := strconv.Atoi(itemIDStr)

		tempItemQuantity := itemClaimReq.ItemQuantity

		tempItemClaimReqs := warehouse.ClaimRequestItem{
			ItemQuantity: itemClaimReq.ItemQuantity,
			MaxClaim:     itemClaimReq.MaxClaim,
			Users:        make([]warehouse.ClaimRequestUser, 0),
		}

		//Lọc qua từng user đã đăng kí nhận đồ trong item
		for idx, user := range itemClaimReq.Users {
			if tempItemQuantity >= user.Quantity {
				userAppointmentItem[user.ID] = append(userAppointmentItem[user.ID], appointment.AppointmentItem{
					ItemID:          uint(itemID), //item id
					ActualQuantity:  int(user.Quantity),
					MissingQuantity: 0,
				})

				tempItemQuantity -= user.Quantity
			} else {
				userAppointmentItem[user.ID] = append(userAppointmentItem[user.ID], appointment.AppointmentItem{
					ItemID:          uint(itemID), //item id
					ActualQuantity:  int(tempItemQuantity),
					MissingQuantity: int(user.Quantity - tempItemQuantity),
				})

				itemClaimReq.Users[idx].Quantity = user.Quantity - tempItemQuantity

				if idx > 0 {
					tempItemClaimReqs.ItemQuantity = tempItemQuantity

					for i := idx; i < len(itemClaimReq.Users); i++ {
						tempItemClaimReqs.Users = append(tempItemClaimReqs.Users, itemClaimReq.Users[idx])
					}
				}

				break
			}
		}

		newItemClaimReqs[key] = tempItemClaimReqs
	}

	for key, value := range userAppointmentItem {
		var (
			startTime time.Time
			endTime   time.Time
		)

		if tempNumPerHour == 0 || currentEndHour >= endAfternoonTime.Hour() {
			appointmentDay, tempNumPerHour, currentStartHour, currentEndHour, _ = c.checkValidDay(ctx, appointmentDay, currentStartHour, currentEndHour, numPerHour, endAfternoonTime, startAfternoonTime, endMorningTime)
		}

		// else if tempNumPerHour == 0 {
		// 	appointmentDay, tempNumPerHour, currentStartHour, currentEndHour, _ = c.checkValidDay(ctx, appointmentDay, currentStartHour, currentEndHour, numPerHour, endAfternoonTime)

		// 	currentStartHour++
		// 	currentEndHour++
		// }

		startTime = time.Date(
			appointmentDay.Year(), appointmentDay.Month(), appointmentDay.Day(), // lấy ngày từ biến `day`
			currentStartHour, 0, 0, 0, // lấy giờ từ biến `hourOnly`
			appointmentDay.Location(), // giữ nguyên location theo `day`
		)

		endTime = time.Date(
			appointmentDay.Year(), appointmentDay.Month(), appointmentDay.Day(), // lấy ngày từ biến `day`
			currentEndHour, 0, 0, 0, // lấy giờ từ biến `hourOnly`
			appointmentDay.Location(), // giữ nguyên location theo `day`
		)

		userAppointment[key] = appointment.Appointment{
			UserID:           key,
			StartTime:        startTime,
			EndTime:          endTime,
			Status:           1,
			AppointmentItems: value,
		}

		tempNumPerHour--
	}

	if len(userAppointment) > 0 {
		if err := c.appointmentRepo.CreateBatch(ctx, userAppointment); err != nil {
			return err
		}

		for key, value := range newItemClaimReqs {
			newItemClaimReqJSON, err := json.Marshal(value)
			if err != nil {
				return err
			}

			//Lưu danh sách người dùng đăng kí mới vào redis
			if err := c.redisRepo.SetToRedisHash(ctx, enums.ItemClaimRequest, key, string(newItemClaimReqJSON)); err != nil {
				return err
			}
		}

		// Lọc qua danh sách các user đã được thông qua đăng kí nhận đồ và cập nhật số lượng đồ
		for key, value := range userAppointmentItem {
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
	}

	return nil
}
