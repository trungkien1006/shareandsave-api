package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"
	"math"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type InterestRepoDB struct {
	db *gorm.DB
}

func NewInterestRepoDB(db *gorm.DB) *InterestRepoDB {
	return &InterestRepoDB{db: db}
}

func (r *InterestRepoDB) GetAll(ctx context.Context, postInterest *[]interest.PostInterest, userID uint, filter interest.GetInterest) (int, error) {
	var (
		query   *gorm.DB
		dbPosts []dbmodel.Post
	)

	if filter.Type == int(enums.InterestTypeInterested) {
		query = r.db.Debug().WithContext(ctx).
			Model(&dbmodel.Post{}).
			Table("post").
			Select("post.id, post.title, post.type, post.slug").
			Preload("Interests").
			Preload("Interests.User").
			Preload("PostItem").
			Preload("PostItem.Item").
			Preload("PostItem.Item.Category").
			Where("author_id = ?", userID).
			Joins("JOIN interest ON interest.post_id = post.id").
			Group("post.id, post.title, post.type, post.slug")
	} else {
		query = r.db.Debug().WithContext(ctx).
			Model(&dbmodel.Post{}).
			Table("post").
			Select("post.id, post.title, post.type, post.slug").
			Preload("Interests").
			Preload("Interests.User").
			Preload("PostItem").
			Preload("PostItem.Item").
			Preload("PostItem.Item.Category").
			Joins("JOIN interest ON interest.post_id = post.id").
			Where("interest.user_id = ?", userID)
	}

	//tim kiem
	if filter.Search != "" {
		query.Where("post.title LIKE ? ", "%"+filter.Search+"%").
			Or("interest.full_name LIKE ?", "%"+filter.Search+"%")
	}

	var totalRecord int64 = 0

	//lay ra tong so record
	if err := query.Count(&totalRecord).Error; err != nil {
		return 0, errors.New("Lỗi khi đếm tổng số record của bài viết: " + err.Error())
	}

	//phan trang
	if filter.Limit != 0 && filter.Page != 0 {
		query.Offset((filter.Page - 1) * filter.Limit).Limit(filter.Limit)
	}

	//sort du lieu
	if filter.Sort != "" && filter.Order != "" {
		filter.Sort = strcase.ToSnake(filter.Sort)

		filter.Sort = "interest." + filter.Sort

		query.Order(filter.Sort + " " + filter.Order)
	}

	if err := query.Find(&dbPosts).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc danh sách quan tâm: " + err.Error())
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(filter.Limit)))

	for _, value := range dbPosts {
		*postInterest = append(*postInterest, dbmodel.GetDTOToDomain(value))
	}

	return totalPage, nil
}

func (r *InterestRepoDB) Create(ctx context.Context, interest interest.Interest) (uint, error) {
	var dbInterest dbmodel.Interest

	dbInterest = dbmodel.CreateDomainToDB(interest)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Interest{}).
		Create(&dbInterest).Error; err != nil {
		return 0, errors.New("Quan tâm không thành công: " + err.Error())
	}

	return dbInterest.ID, nil
}

func (r *InterestRepoDB) Delete(ctx context.Context, postID uint, userID uint) (uint, error) {
	var (
		count      int64
		dbInterest dbmodel.Interest
	)

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Interest{}).Where("user_id = ? AND post_id = ?", userID, postID).Find(&dbInterest).Error; err != nil {
		return 0, errors.New("Có lỗi khi tìm kiếm quan tâm: " + err.Error())
	}

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Transaction{}).Where("interest_id = ?", dbInterest.ID).Count(&count).Error; err != nil {
		return 0, errors.New("Lỗi khi kiểm tra giao dịch trong quan tâm: " + err.Error())
	}

	if count > 0 {
		return 0, errors.New("Không thể hủy quan tâm do đã phát sinh giao dịch")
	}

	count = 0

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Comment{}).Where("interest_id = ?", dbInterest.ID).Count(&count).Error; err != nil {
		return 0, errors.New("Lỗi khi kiểm tra tin nhắn trong quan tâm: " + err.Error())
	}

	if count > 0 {
		return 0, errors.New("Không thể hủy quan tâm do đã có tin nhắn")
	}

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Interest{}).Delete(&dbInterest).Error; err != nil {
		return 0, errors.New("Lỗi khi xóa quan tâm: " + err.Error())
	}

	return dbInterest.ID, nil
}

func (r *InterestRepoDB) IsExist(ctx context.Context, userID uint, postID uint) (bool, error) {
	var count int64 = 0

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Interest{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count).Error; err != nil {
		return false, errors.New("Có lỗi khi kiểm tra quan tâm tồn tại: " + err.Error())
	}

	return count > 0, nil
}
