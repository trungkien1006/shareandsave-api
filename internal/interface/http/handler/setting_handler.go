package handler

import (
	"final_project/internal/application/app/settingapp"
	"final_project/internal/domain/setting"
	settingdto "final_project/internal/dto/settingDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SettingHandler struct {
	uc *settingapp.UseCase
}

func NewSettingHandler(uc *settingapp.UseCase) *SettingHandler {
	return &SettingHandler{uc: uc}
}

// @Summary Get settings
// @Description API lấy ra tất cả cấu hình
// @Security BearerAuth
// @Tags settings
// @Accept json
// @Produce json
// @Success 200 {object} settingdto.GetSettingResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /settings [get]
func (h *SettingHandler) GetAll(c *gin.Context) {
	var (
		settings      []setting.Setting
		settingDTORes []settingdto.SettingDTO
	)

	if err := h.uc.GetAllSetting(c.Request.Context(), &settings); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	for _, value := range settings {
		settingDTORes = append(settingDTORes, settingdto.SettingDomainToDTO(value))
	}

	c.JSON(http.StatusOK, settingdto.GetSettingResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched settings successfully",
		Data: settingdto.GetSettingResponse{
			Settings: settingDTORes,
		},
	})
}

// @Summary Get settings
// @Description API lấy ra cấu hình theo key
// @Security BearerAuth
// @Tags settings
// @Accept json
// @Produce json
// @Param settingKey path string true "Key setting"
// @Success 200 {object} settingdto.GetSettingByKeyResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /settings/{settingKey} [get]
func (h *SettingHandler) GetByKey(c *gin.Context) {
	var (
		req           settingdto.GetSettingByKeyRequest
		setting       setting.Setting
		settingDTORes settingdto.SettingDTO
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.GetSettingByKey(c.Request.Context(), &setting, req.SettingKey); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	settingDTORes = settingdto.SettingDomainToDTO(setting)

	c.JSON(http.StatusOK, settingdto.GetSettingByKeyResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched setting by key successfully",
		Data: settingdto.GetSettingByKeyResponse{
			Setting: settingDTORes,
		},
	})
}

// @Summary Update settings
// @Description API cập nhật giá trị cấu hình
// @Security BearerAuth
// @Tags settings
// @Accept json
// @Produce json
// @Param settingKey path string true "Key setting"
// @Param request body settingdto.UpdateSettingRequest true "Update setting info"
// @Success 200 {object} settingdto.GetSettingByKeyResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /settings/{settingKey} [patch]
func (h *SettingHandler) Update(c *gin.Context) {
	var (
		req     settingdto.UpdateSettingRequest
		setting setting.Setting
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	settingKey := c.Param("settingKey")

	setting.Key = settingKey
	setting.Value = req.Value

	if err := h.uc.UpdateSetting(c.Request.Context(), &setting, settingKey); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, settingdto.UpdateSettingResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated setting by key successfully",
		Data:    gin.H{},
	})
}
