package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/interest"
	"final_project/internal/infrastructure/persistence/dbmodel"
	"final_project/internal/pkg/enums"
	"math"
	"sort"

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
			Select("post.id, post.title, post.type, post.slug, post.author_id, post.updated_at, post.created_at, post.description").
			Preload("Interests", func(db *gorm.DB) *gorm.DB {
				return db.Where("user_id = ?", userID)
			}).
			Preload("Interests.Comments", func(db *gorm.DB) *gorm.DB {
				// Lấy comment chưa đọc và sort theo thời gian tạo mới nhất
				return db.Where("is_read = 0 AND sender_id != ?", userID).Order("created_at DESC")
			}).
			Preload("Interests.NewComment", func(db *gorm.DB) *gorm.DB {
				// Lấy comment chưa đọc và sort theo thời gian tạo mới nhất
				return db.Where("sender_id != ?", userID).Order("created_at DESC").Offset(0).Limit(1)
			}).
			Preload("Author").
			Preload("Interests.User").
			Preload("PostItem").
			Preload("PostItem.Item").
			Preload("PostItem.Item.Category").
			Where("interest.user_id = ? AND interest.deleted_at IS NULL", userID).
			Joins("JOIN interest ON interest.post_id = post.id").
			Joins("JOIN user ON interest.user_id = user.id").
			Group("post.id, post.title, post.type, post.slug, post.author_id, post.updated_at, post.created_at, post.description")
	} else {
		query = r.db.Debug().WithContext(ctx).
			Model(&dbmodel.Post{}).
			Table("post").
			Select("post.id, post.title, post.type, post.slug, post.author_id, post.updated_at, post.created_at, post.description").
			Preload("Author").
			Preload("Interests", func(db *gorm.DB) *gorm.DB {
				return db.Where("deleted_at IS NULL")
			}).
			Preload("Interests.Comments", func(db *gorm.DB) *gorm.DB {
				// Lấy comment chưa đọc và sort theo thời gian tạo mới nhất
				return db.Where("is_read = 0 AND sender_id != ?", userID).Order("created_at DESC")
			}).
			Preload("Interests.NewComment", func(db *gorm.DB) *gorm.DB {
				// Lấy comment chưa đọc và sort theo thời gian tạo mới nhất
				return db.Where("sender_id != ?", userID).Order("created_at DESC").Offset(0).Limit(1)
			}).
			// Preload().
			Preload("Interests.User").
			Preload("PostItem").
			Preload("PostItem.Item").
			Preload("PostItem.Item.Category").
			Joins("JOIN interest ON interest.post_id = post.id AND interest.deleted_at IS NULL").
			Joins("JOIN user ON interest.user_id = user.id").
			Where("post.author_id = ? AND post.deleted_at IS NULL", userID).
			Group("post.id, post.title, post.type, post.slug, post.author_id, post.updated_at, post.created_at, post.description")
	}

	//tim kiem
	if filter.Search != "" {
		query.Where("( post.title LIKE ? OR user.full_name LIKE ? )", "%"+filter.Search+"%", "%"+filter.Search+"%")
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

		if filter.Sort == "interest.created_at" {
			filter.Sort = "post.updated_at"
		}

		query.Order(filter.Sort + " " + filter.Order)
	}

	if err := query.Find(&dbPosts).Error; err != nil {
		return 0, errors.New("Lỗi khi lọc danh sách quan tâm: " + err.Error())
	}

	for i := range dbPosts {
		sort.SliceStable(dbPosts[i].Interests, func(a, b int) bool {
			aComments := dbPosts[i].Interests[a].Comments
			bComments := dbPosts[i].Interests[b].Comments

			// Nếu cả hai đều có comment thì so sánh created_at mới nhất
			if len(aComments) > 0 && len(bComments) > 0 {
				return aComments[0].CreatedAt.After(bComments[0].CreatedAt)
			}
			// Ưu tiên interest có comment hơn
			return len(aComments) > len(bComments)
		})
	}

	//tinh toan total page
	totalPage := int(math.Ceil(float64(totalRecord) / float64(filter.Limit)))

	for _, value := range dbPosts {
		*postInterest = append(*postInterest, dbmodel.GetDTOToDomain(value))
	}

	return totalPage, nil
}

func (r *InterestRepoDB) GetTotalUnreadMessage(ctx context.Context, userID uint, interestType enums.InterestType) (uint, error) {
	var (
		unreadCount int64
		query       *gorm.DB
	)

	query = r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Comment{}).
		Table("comment").
		Where("comment.deleted_at IS NULL AND comment.sender_id != ? AND comment.receiver_id = ? AND comment.is_read = 0", userID, userID).
		Joins("JOIN interest ON interest.id = comment.interest_id AND interest.deleted_at IS NULL")

	if interestType == enums.InterestTypeInterested {
		query.Where("interest.user_id = ?", userID)
	} else if interestType == enums.InterestTypeFollowedBy {
		query.Where("interest.user_id != ?", userID)
	}

	if err := query.Count(&unreadCount).Error; err != nil {
		return 0, errors.New("Có lỗi khi đếm số tin nhắn chưa đọc: " + err.Error())
	}

	return uint(unreadCount), nil
}

func (r *InterestRepoDB) GetDetailByID(ctx context.Context, postInterest *interest.PostInterest, interestID uint) error {
	var (
		dbPosts dbmodel.Post
	)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Post{}).
		Table("post").
		Select("post.id, post.title, post.type, post.slug, post.author_id, post.updated_at, post.created_at, post.description").
		Preload("Interests", func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", interestID)
		}).
		Preload("Author").
		Preload("Interests.User").
		Preload("PostItem").
		Preload("PostItem.Item").
		Preload("PostItem.Item.Category").
		Where("interest.id = ? AND interest.deleted_at IS NULL", interestID).
		Joins("JOIN interest ON interest.post_id = post.id").
		Joins("JOIN user ON interest.user_id = user.id").
		Group("post.id, post.title, post.type, post.slug, post.author_id, post.updated_at, post.created_at, post.description").
		Find(&dbPosts).Error; err != nil {
		return errors.New("Lỗi khi lọc danh sách quan tâm: " + err.Error())
	}

	*postInterest = dbmodel.GetDTOToDomain(dbPosts)

	return nil
}

func (r *InterestRepoDB) Create(ctx context.Context, interest interest.Interest) (uint, error) {
	var (
		dbInterest dbmodel.Interest
		authorID   uint = 0
	)

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Post{}).Select("author_id").Where("id = ? AND status != 2", interest.PostID).Scan(&authorID).Error; err != nil {
		return 0, errors.New("Có lỗi khi kiểm tra quan tâm chính mình: " + err.Error())
	}

	if authorID == 0 {
		return 0, errors.New("Không thể quan tâm bài viết đã từ chối:")
	}

	if authorID == interest.UserID {
		return 0, errors.New("Không thể quan tâm chính bài viết của mình:")
	}

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

	if err := r.db.Debug().WithContext(ctx).Delete(&dbmodel.Interest{}, dbInterest.ID).Error; err != nil {
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

func (r *InterestRepoDB) IsExistByID(ctx context.Context, interestID uint) (bool, error) {
	var count int64 = 0

	if err := r.db.Debug().WithContext(ctx).Model(&dbmodel.Interest{}).Where("id = ?", interestID).Count(&count).Error; err != nil {
		return false, errors.New("Có lỗi khi kiểm tra quan tâm tồn tại: " + err.Error())
	}

	return count > 0, nil
}
