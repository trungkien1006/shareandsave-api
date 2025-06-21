package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/setting"
	"final_project/internal/infrastructure/persistence/dbmodel"

	"gorm.io/gorm"
)

type SettingRepoDB struct {
	db *gorm.DB
}

func NewSettingRepoDB(db *gorm.DB) *SettingRepoDB {
	return &SettingRepoDB{db: db}
}

func (r *SettingRepoDB) GetAll(ctx context.Context, settings *[]setting.Setting) error {
	var dbSettings []dbmodel.Setting

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Setting{}).
		Find(&dbSettings).Error; err != nil {
		return errors.New("Có lỗi khi truy suất danh sách cấu hình: " + err.Error())
	}

	for _, value := range dbSettings {
		*settings = append(*settings, dbmodel.SettingDBToDomain(value))
	}

	return nil
}

func (r *SettingRepoDB) GetByKey(ctx context.Context, setting *setting.Setting, settingKey string) error {
	var dbSetting dbmodel.Setting

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Setting{}).
		Where("key = ?", settingKey).
		First(&dbSetting).Error; err != nil {
		return errors.New("Có lỗi khi truy xuất cấu hình theo key: " + err.Error())
	}

	*setting = dbmodel.SettingDBToDomain(dbSetting)

	return nil
}

func (r *SettingRepoDB) GetByID(ctx context.Context, setting *setting.Setting, settingID uint) error {
	var dbSetting dbmodel.Setting

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Setting{}).
		Where("id = ?", settingID).
		First(&dbSetting).Error; err != nil {
		return errors.New("Có lỗi khi truy xuất cấu hình theo id: " + err.Error())
	}

	*setting = dbmodel.SettingDBToDomain(dbSetting)

	return nil
}

func (r *SettingRepoDB) Create(ctx context.Context, setting setting.Setting) error {
	dbSetting := dbmodel.SettingDomainToDB(setting)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Setting{}).
		Create(&dbSetting).Error; err != nil {
		return errors.New("Có lỗi khi thêm setting: " + err.Error())
	}

	return nil
}

func (r *SettingRepoDB) Update(ctx context.Context, updateSetting setting.Setting) error {
	dbSetting := dbmodel.SettingDomainToDB(updateSetting)

	if err := r.db.Debug().WithContext(ctx).
		Model(&dbmodel.Setting{}).
		Where("key = ?", dbSetting.Key).
		Update("value", dbSetting.Value).Error; err != nil {
		return errors.New("Có lỗi khi cập nhật cấu hình: " + err.Error())
	}

	return nil
}
