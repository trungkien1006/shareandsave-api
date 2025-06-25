package persistence

import (
	"gorm.io/gorm"
)

type LeaveRequestsRepoDB struct {
	db *gorm.DB
}

func NewLeaveRequestsRepoDB(db *gorm.DB) *LeaveRequestsRepoDB {
	return &LeaveRequestsRepoDB{db: db}
}

// func (r *LeaveRequestsRepoDB) GetAll(ctx context.Context, leaveRequests *[]leaverequests.LeaveRequest, req filter.FilterRequest) (int, error) {
// 	var (
// 		query           *gorm.DB
// 		totalRecords    int64
// 		dbLeaveRequests []dbmodel.LeaveRequests
// 	)

// 	query = r.db.Debug().
// 		WithContext(ctx).
// 		Model(&dbmodel.LeaveRequests{}).
// 		Table("leave_requests as lr").
// 		Joins("JOIN user ON user.id = lr.user_id").
// 		Preload("User")

// 	if req.SearchBy != "" && req.SearchValue != "" {
// 		column := strcase.ToSnake(req.SearchBy) // "fullName" -> "full_name"

// 		if column == "user_name" {
// 			column = "user.full_name"
// 		} else {
// 			column = "lr." + column
// 		}

// 		query.Where(column+" LIKE ? ", "%"+req.SearchValue+"%")
// 	}

// 	if err := query.Count(&totalRecords).Error; err != nil {
// 		return 0, errors.New("Có lỗi khi truy vấn danh sách phiếu nghỉ phép: " + err.Error())
// 	}

// 	if req.Sort != "" && req.Order != "" {
// 		query = query.Order("lr." + strcase.ToSnake(req.Sort) + " " + req.Order)
// 	}

// 	if req.Limit > 0 && req.Page > 0 {
// 		query.Offset((req.Page - 1) * req.Limit).Limit(req.Limit)
// 	}

// 	if err := query.Find(&dbLeaveRequests).Error; err != nil {
// 		return 0, errors.New("Có lỗi khi truy vấn danh sách phiếu nghỉ phép: " + err.Error())
// 	}

// 	totalPages := int((totalRecords + int64(req.Limit) - 1) / int64(req.Limit))

// 	for _, value := range dbLeaveRequests {
// 		*leaveRequests = append(*leaveRequests, dbmodel.LeaveReqDBToDomain(value))
// 	}

// 	return totalPages, nil
// }

// func (r *LeaveRequestsRepoDB) GetByID(ctx context.Context, leaveRequest *leaverequests.LeaveRequest, leaveReqID uint) error {
// 	var dbLeaveReq dbmodel.LeaveRequests

// 	if err := r.db.Debug().
// 		WithContext(ctx).
// 		Model(&dbmodel.LeaveRequests{}).
// 		Table("leave_requests as lr").
// 		Where("lr.id = ?", leaveReqID).
// 		Joins("JOIN user ON user.id = lr.user_id").
// 		Preload("User").
// 		First(&dbLeaveReq).Error; err != nil {
// 		return errors.New("Có lỗi khi truy xuất phiếu nghỉ phép bằng id: " + err.Error())
// 	}

// 	*leaveRequest = dbmodel.LeaveReqDBToDomain(dbLeaveReq)

// 	return nil
// }

// func (r *LeaveRequestsRepoDB) Create(ctx context.Context, leaveReq leaverequests.LeaveRequest) error {
// 	dbLeaveRequest := dbmodel.LeaveReqDomainToDB(leaveReq)

// 	if err := r.db.Debug().WithContext(ctx).
// 		Model(&dbmodel.LeaveRequests{}).
// 		Create(&dbLeaveRequest).Error; err != nil {
// 		return errors.New("Có lỗi khi tạo")
// 	}

// 	return nil
// }

// func (r *LeaveRequestsRepoDB) IsInLeaveRequest(ctx context.Context, day time.Time) (bool, error) {
// 	var count int64
// 	err := r.db.WithContext(ctx).
// 		Model(&dbmodel.LeaveRequests{}).
// 		Where("start_date <= ? AND end_date >= ?", day, day).
// 		Count(&count).Error

// 	if err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }
