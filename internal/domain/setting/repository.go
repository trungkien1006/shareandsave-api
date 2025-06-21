package setting

import "context"

type Repository interface {
	GetAll(ctx context.Context, setting *[]Setting) error
	GetByID(ctx context.Context, setting *Setting, settingID uint) error
	GetByKey(ctx context.Context, setting *Setting, settingKey string) error
	Create(ctx context.Context, setting Setting) error
	Update(ctx context.Context, updateSetting Setting) error
}
