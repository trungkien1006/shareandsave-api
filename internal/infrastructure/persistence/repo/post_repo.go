package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/post"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"math"

	"github.com/iancoleman/strcase"
	"gorm.io/gorm"
)

type PostRepoDB struct {
	db *gorm.DB
}

func NewPostRepoDB(db *gorm.DB) *PostRepoDB {
	return &PostRepoDB{db: db}
}

func (r *PostRepoDB) AdminGetAll(ctx context.Context, posts *[]post.Post, filter post.PostFilterRequest) (int, error) {
	var (
		query  *gorm.DB
		dbPost []dbmodel.Post
	)

	query = r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Post{}).
		Table("post as post").
		Joins("JOIN user AS author ON author.id = post.author_id").
		Preload("Author")

	if filter.SearchBy != "" && filter.SearchValue != "" {
		column := strcase.ToSnake(filter.SearchBy) // "fullName" -> "full_name"

		if column == "author_name" {
			column = "author.full_name"
		} else {
			column = "post." + column
		}

		query.Where(column+" LIKE ? ", "%"+filter.SearchValue+"%")

	}

	if filter.Type != 0 {
		query.Where("post.type = ?", filter.Type)
	}

	if filter.Status != 0 {
		query.Where("post.status = ?", filter.Status)
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
	if filter.Sort != "" {
		if filter.Sort == "authorName" {
			filter.Sort = "author.full_name"
		} else {
			filter.Sort = strcase.ToSnake(filter.Sort)
		}

		query.Order(filter.Sort + " " + filter.Order)
	}

	if err := query.Find(&dbPost).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc bài viết: " + err.Error())
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(filter.Limit)))

	for _, value := range dbPost {
		*posts = append(*posts, dbmodel.AdminPostDBToDomain(value))
	}

	return totalPage, nil
}

func (r *PostRepoDB) GetDetailByID(ctx context.Context, post *post.DetailPost, postID uint) error {
	var dbPost dbmodel.Post

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Post{}).
		Where("id = ?", postID).
		Preload("Author").
		Preload("Interests").
		Preload("Interests.User").
		Preload("PostItem").
		Preload("PostItem.Item").
		Preload("PostItem.Item.Category").
		Find(&dbPost).Error; err != nil {
		return errors.New("Có lỗi khi tìm kiếm bài viết theo ID: " + err.Error())
	}

	*post = dbmodel.DetailPostDBToDomain(dbPost)

	return nil
}

func (r *PostRepoDB) GetByID(ctx context.Context, post *post.Post, postID uint) error {
	var dbPost dbmodel.Post

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Post{}).
		Where("id = ?", postID).
		Preload("Author").
		Find(&dbPost).Error; err != nil {
		return errors.New("Có lỗi khi tìm kiếm bài viết theo ID: " + err.Error())
	}

	*post = dbmodel.AdminPostDBToDomain(dbPost)

	return nil
}

func (r *PostRepoDB) Save(ctx context.Context, post *post.CreatePost) error {
	tx := r.db.Begin()

	dbPost := dbmodel.CreatePostDomainToDB(*post)

	if err := tx.Debug().WithContext(ctx).Model(&dbmodel.Post{}).Create(&dbPost).Error; err != nil {
		return errors.New("Lỗi khi tạo bài đăng: " + err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("Có lỗi khi commit transaction: " + err.Error())
	}

	return nil
}

func (r *PostRepoDB) Update(ctx context.Context, post *post.Post) error {
	dbPost := dbmodel.AdminPostDomainToDB(*post)

	if err := r.db.Debug().
		WithContext(ctx).
		Model(&dbmodel.Post{}).
		Omit("CreatedAt").
		Omit("DeleteAt").
		Where("id = ?", post.ID).
		Updates(&dbPost).Error; err != nil {
		return errors.New("Lỗi khi cập nhật bài viết: " + err.Error())
	}

	return nil
}
