package settingdto

type GetSettingResponse struct {
	Settings []SettingDTO `json:"settings"`
}

type GetSettingResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetSettingResponse `json:"data"`
}

type GetSettingByKeyResponse struct {
	Setting SettingDTO `json:"setting"`
}

type GetSettingByKeyResponseWrapper struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetSettingByKeyResponse `json:"data"`
}

type UpdateSettingResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
