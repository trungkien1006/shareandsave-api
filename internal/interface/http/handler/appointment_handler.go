package handler

import (
	"final_project/internal/application/app/appointmentapp"
	"final_project/internal/domain/appointment"
	appointmentdto "final_project/internal/dto/appointmentDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"

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
