package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/infrastructure/persistence/dbmodel"
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

	if filter.Type == 1 {
		query = r.db.Debug().WithContext(ctx).
			Model(&dbmodel.Post{}).
			Table("post").
			Select("post.id, post.title, post.type").
			Preload("Interests").
			Preload("Interests.User").
			Preload("PostItem").
			Preload("PostItem.Item").
			Preload("PostItem.Item.Category").
			Where("author_id = ?", userID).
			Joins("JOIN interest ON interest.post_id = post.id").
			Group("post.id, post.title, post.type")
	} else {
		query = r.db.Debug().WithContext(ctx).
			Model(&dbmodel.Post{}).
			Table("post").
			Select("post.id, post.title, post.type").
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

func (r *InterestRepoDB) Create(ctx context.Context, interest interest.Interest) error {
	var dbInterest dbmodel.Interest

	dbInterest = dbmodel.CreateDomainToDB(interest)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Interest{}).
		Create(&dbInterest).Error; err != nil {
		return errors.New("Quan tâm không thành công: " + err.Error())
	}

	return nil
}
