package handler

import (
	"final_project/internal/application/app/appointmentapp"
	"final_project/internal/domain/appointment"
	appointmentdto "final_project/internal/dto/appointmentDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	uc *appointmentapp.UseCase
}

func NewAppointmentHandler(uc *appointmentapp.UseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
}

// @Summary Get appointments
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags appointments
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (startTime, endTime)"
// @Param order query string false "Sort type" enum(ASC,DESC) example(ASC, DESC)
// @Param   searchBy   query    string  false  "Trường lọc (vd: status, userName)"
// @Param   searchValue   query    string  false  "Giá trị lọc"
// @Success 200 {object} appointmentdto.GetAppointmentResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /appointments [get]
func (h *AppointmentHandler) GetAll(c *gin.Context) {
	var (
		req                appointmentdto.GetAllAppointmentRequest
		domainAppointments []appointment.Appointment
		domainFilter       appointment.FilterAllAppointment
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	req.SetDefault()

	domainFilter.Page = req.Page
	domainFilter.Limit = req.Limit
	domainFilter.Sort = req.Sort
	domainFilter.Order = req.Order
	domainFilter.SearchBy = req.SearchBy
	domainFilter.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAll(c.Request.Context(), &domainAppointments, domainFilter, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	appointmentDTORes := make([]appointmentdto.AppointmentDTO, 0)

	for _, value := range domainAppointments {
		appointmentDTORes = append(appointmentDTORes, appointmentdto.AppointmentDomainToDTO(value))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched appointments successfully",
		"data": appointmentdto.GetAppointmentResponse{
			Appointments: appointmentDTORes,
			TotalPage:    totalPage,
		},
	})
}

// @Summary Get my appointments
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags appointments
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (startTime, endTime)"
// @Param order query string false "Sort type" enum(ASC,DESC) example(ASC, DESC)
// @Param   searchBy   query    string  false  "Trường lọc (vd: status)"
// @Param   searchValue   query    string  false  "Giá trị lọc"
// @Success 200 {object} appointmentdto.GetAppointmentResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /client/appointments [get]
func (h *AppointmentHandler) ClientGetAll(c *gin.Context) {
	var (
		req                appointmentdto.GetAllMyAppointmentRequest
		domainAppointments []appointment.Appointment
		domainFilter       appointment.FilterAllAppointment
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	req.SetDefault()

	domainFilter.Page = req.Page
	domainFilter.Limit = req.Limit
	domainFilter.Sort = req.Sort
	domainFilter.Order = req.Order
	domainFilter.SearchBy = req.SearchBy
	domainFilter.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAll(c.Request.Context(), &domainAppointments, domainFilter, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	appointmentDTORes := make([]appointmentdto.AppointmentDTO, 0)

	for _, value := range domainAppointments {
		appointmentDTORes = append(appointmentDTORes, appointmentdto.AppointmentDomainToDTO(value))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched client appointments successfully",
		"data": appointmentdto.GetAppointmentResponse{
			Appointments: appointmentDTORes,
			TotalPage:    totalPage,
		},
	})
}

// @Summary Get appointment by ID
// @Description API lấy thông tin item theo ID
// @Security BearerAuth
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointmentID path int true "ID appointment"
// @Success 200 {object} appointmentdto.GetAppointmentByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /appointments/{appointmentID} [get]
func (h *AppointmentHandler) GetByID(c *gin.Context) {
	var (
		req               appointmentdto.GetAppointmentByIDRequest
		domainAppointment appointment.Appointment
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := h.uc.GetByID(c.Request.Context(), &domainAppointment, req.AppointmentID); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	appointmentDTORes := appointmentdto.AppointmentDomainToDTO(domainAppointment)

	c.JSON(http.StatusOK, appointmentdto.GetAppointmentByIDResponseWrapper{
		Code:    http.StatusOK,
		Message: "Get appointment by id successfully",
		Data: appointmentdto.GetAppointmentByIDResponse{
			Appointment: appointmentDTORes,
		},
	})
}

// @Summary Update appointment
// @Description API cập nhật appointment
// @Security BearerAuth
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointmentID path int true "ID appointment"
// @Param request body appointmentdto.UpdateAppointmentRequest true "Update appointment info"
// @Success 200 {object} appointmentdto.UpdateAppointmentResponseWrapper "Updated appointment successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /appointments/{appointmentID} [patch]
func (h *AppointmentHandler) Update(c *gin.Context) {
	var (
		req               appointmentdto.UpdateAppointmentRequest
		domainAppointment appointment.Appointment
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	appointmentID, err := strconv.Atoi(c.Param("appointmentID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	domainAppointment = appointmentdto.UpdateAppointmentDTOToDomain(req)

	if err := h.uc.Update(c.Request.Context(), domainAppointment, uint(appointmentID)); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	c.JSON(http.StatusOK, appointmentdto.UpdateAppointmentResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated appointment successfully",
		Data:    gin.H{},
	})
}

// @Summary Update batch appointment
// @Description API cập nhật danh sách appointment
// @Security BearerAuth
// @Tags appointments
// @Accept json
// @Produce json
// @Param request body appointmentdto.UpdateBatchAppointmentRequest true "Update batch appointment info"
// @Success 200 {object} appointmentdto.UpdateAppointmentResponseWrapper "Updated batch appointment successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /appointments [patch]
func (h *AppointmentHandler) UpdateBatch(c *gin.Context) {
	var (
		req               appointmentdto.UpdateBatchAppointmentRequest
		domainAppointment []appointment.Appointment
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	for _, value := range req.Appointments {
		domainAppointment = append(domainAppointment, appointmentdto.UpdateAppointmentDTOToDomain(value))
	}

	if err := h.uc.UpdateBatch(c.Request.Context(), domainAppointment, req.AppointmentIDs); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	c.JSON(http.StatusOK, appointmentdto.UpdateAppointmentResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated batch appointment successfully",
		Data:    gin.H{},
	})
}
