package settingdto

import "final_project/internal/domain/setting"

// Domain to DTO
func SettingDomainToDTO(domain setting.Setting) SettingDTO {
	return SettingDTO{
		ID:    domain.ID,
		Key:   domain.Key,
		Value: domain.Value,
	}
}
