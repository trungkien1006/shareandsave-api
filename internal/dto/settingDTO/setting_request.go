package settingdto

type GetSettingByKeyRequest struct {
	SettingKey string `uri:"settingKey"`
}

type UpdateSettingRequest struct {
	Value string `json:"value"`
}
