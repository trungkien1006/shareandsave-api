package handler

import (
	"final_project/internal/application/warehouseapp"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/warehouse"
	warehousedto "final_project/internal/dto/warehouseDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	uc *warehouseapp.UseCase
}

func NewWarehouseHandler(uc *warehouseapp.UseCase) *WarehouseHandler {
	return &WarehouseHandler{uc: uc}
}

// @Summary Get warehouse
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags warehouses
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Param sort query string false "Sort column (createdAt quantity)"
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   searchBy   query    string  false  "Trường lọc (senderName invoiceNum itemName classify sku description stockPlace)"
// @Param   searchValue   query    string  false  "Giá trị lọc:"
// @Success 200 {object} warehousedto.FilterWarehouseResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /warehouses [get]
func (h *WarehouseHandler) GetAll(c *gin.Context) {
	var (
		req             warehousedto.GetWarehouseRequest
		domainFilter    filter.FilterRequest
		domainWarehouse []warehouse.Warehouse
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	req.SetDefault()

	domainFilter.Page = req.Page
	domainFilter.Limit = req.Limit
	domainFilter.Sort = req.Sort
	domainFilter.Order = req.Order
	domainFilter.SearchBy = req.SearchBy
	domainFilter.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAllWarehouse(c.Request.Context(), &domainWarehouse, domainFilter)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	warehouseDTORes := make([]warehousedto.WarehouseDTO, 0)

	for _, value := range domainWarehouse {
		warehouseDTORes = append(warehouseDTORes, warehousedto.WarehouseDomainToDTO(value))
	}

	c.JSON(http.StatusOK, warehousedto.FilterWarehouseResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched warehouses successfully",
		Data: warehousedto.FilterWarehouseResponse{
			Warehouses: warehouseDTORes,
			TotalPage:  totalPage,
		},
	})
}

// @Summary Get warehouse by ID
// @Description API lấy thông tin warehouse theo ID
// @Tags warehouses
// @Accept json
// @Produce json
// @Param warehouseID path int true "ID warehouse"
// @Success 200 {object} warehousedto.GetWarehouseByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /warehouses/{warehouseID} [get]
func (h *WarehouseHandler) GetByID(c *gin.Context) {
	var (
		req             warehousedto.GetWarehouseByIDRequest
		domainWarehouse warehouse.DetailWarehouse
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.GetWarehouseByID(c.Request.Context(), &domainWarehouse, req.WarehouseID); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	warehouseDTORes := warehousedto.DetailWarehouseDomainToDTO(domainWarehouse)

	c.JSON(http.StatusOK, warehousedto.GetWarehouseByIDResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched warehouses successfully",
		Data: warehousedto.GetWarehouseByIDResponse{
			Warehouse: warehouseDTORes,
		},
	})
}
