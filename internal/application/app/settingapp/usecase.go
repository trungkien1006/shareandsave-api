package settingapp

import (
	"context"
	"final_project/internal/domain/setting"
)

type UseCase struct {
	repo setting.Repository
}

func NewUseCase(r setting.Repository) *UseCase {
	return &UseCase{repo: r}
}

func (uc *UseCase) GetAllSetting(ctx context.Context, settings *[]setting.Setting) error {
	if err := uc.repo.GetAll(ctx, settings); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetSettingByKey(ctx context.Context, setting *setting.Setting, settingKey string) error {
	if err := uc.repo.GetByKey(ctx, setting, settingKey); err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) UpdateSetting(ctx context.Context, domainSetting *setting.Setting, settingKey string) error {
	var updateSetting setting.Setting

	if err := uc.repo.GetByKey(ctx, &updateSetting, settingKey); err != nil {
		return err
	}

	updateSetting.Value = domainSetting.Value

	if err := uc.repo.Update(ctx, updateSetting); err != nil {
		return err
	}

	return nil
}
