package handler

import (
	"final_project/internal/application/app/warehouseapp"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/warehouse"
	warehousedto "final_project/internal/dto/warehouseDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"
	"strconv"

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
// @Param   searchBy   query    string  false  "Trường lọc (senderName receiverName invoiceNum itemName classify sku description stockPlace)"
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

// @Summary Get item warehouse
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags item warehouses
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Param sort query string false "Sort column (createdAt)"
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   searchBy   query    string  false  "Trường lọc (itemName description code status)"
// @Param   searchValue   query    string  false  "Giá trị lọc:"
// @Success 200 {object} warehousedto.FilterItemWarehouseResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Security BearerAuth
// @Router /item-warehouses [get]
func (h *WarehouseHandler) GetAllItem(c *gin.Context) {
	var (
		req                 warehousedto.GetItemWarehouseRequest
		domainFilter        filter.FilterRequest
		domainItemWarehouse []warehouse.ItemWareHouse
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

	totalPage, err := h.uc.GetAllItemWarehouse(c.Request.Context(), &domainItemWarehouse, domainFilter)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	itemWarehouseDTORes := make([]warehousedto.ItemWareHouseDTO, 0)

	for _, value := range domainItemWarehouse {
		itemWarehouseDTORes = append(itemWarehouseDTORes, warehousedto.ItemWarehouseDomainToDTO(value))
	}

	c.JSON(http.StatusOK, warehousedto.FilterItemWarehouseResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched item warehouses successfully",
		Data: warehousedto.FilterItemWarehouseResponse{
			ItemWarehouses: itemWarehouseDTORes,
			TotalPage:      totalPage,
		},
	})
}

// @Summary Get item old stock
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags item warehouses
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Param sort query string false "Sort column (quantity)"
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   categoryID   query    int  false  "Phân loại đồ"
// @Param   search   query    string  false  "Lọc"
// @Success 200 {object} warehousedto.FilterItemOldStockResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/item-warehouses/old-stock [get]
func (h *WarehouseHandler) GetAllItemOldStock(c *gin.Context) {
	var (
		req          warehousedto.GetItemOldStockRequest
		domainItems  []warehouse.ItemOldStock
		domainFilter warehouse.GetItemOldStockFilter
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
	domainFilter.Search = req.Search
	domainFilter.CategoryID = req.CategoryID

	claimRequestCounts, totalPage, err := h.uc.GetAllItemOldStock(c.Request.Context(), &domainItems, domainFilter)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	for key, value := range claimRequestCounts {
		domainItems[key].ClaimItemRequests = value
	}

	itemOldStockDTORes := make([]warehousedto.ItemOldStockDTO, 0)

	for _, value := range domainItems {
		itemOldStockDTORes = append(itemOldStockDTORes, warehousedto.ItemOldStockDomainToDTO(value))
	}

	c.JSON(http.StatusOK, warehousedto.FilterItemOldStockResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched item old stock successfully",
		Data: warehousedto.FilterItemOldStockResponse{
			ItemOldStocks: itemOldStockDTORes,
			TotalPage:     totalPage,
		},
	})
}

// @Summary Get my claim requests
// @Description API lấy danh sách các món đồ đã đăng kí nhận
// @Tags item warehouses
// @Accept json
// @Produce json
// @Success 200 {object} warehousedto.GetMyClaimRequestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Security BearerAuth
// @Router /client/item-warehouses/claim-request [get]
func (h *WarehouseHandler) GetMyClaimRequest(c *gin.Context) {
	domainUserClaimReq := make([]warehouse.MyClaimRequest, 0)

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := h.uc.GetMyClaimRequest(c.Request.Context(), &domainUserClaimReq, userID); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	userClaimsRequestDTORes := make([]warehousedto.MyClaimRequestDTO, 0)

	for _, value := range domainUserClaimReq {
		userClaimsRequestDTORes = append(userClaimsRequestDTORes, warehousedto.MyClaimRequestDomainToDTO(value))
	}

	c.JSON(http.StatusOK, warehousedto.GetMyClaimRequestResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched claim requests successfully",
		Data: warehousedto.GetMyClaimRequestResponse{
			ClaimRequests: userClaimsRequestDTORes,
		},
	})
}

// @Summary Create claim request
// @Description API lưu thông tin đăng kí nhận đồ
// @Security BearerAuth
// @Tags item warehouses
// @Accept json
// @Produce json
// @Param request body []warehousedto.CreateClaimRequestRequest true "Claim request creation payload"
// @Success 201 {object} warehousedto.GetItemWarehouseByCodeResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/item-warehouses/claim-request [post]
func (h *WarehouseHandler) CreateClaimRequest(c *gin.Context) {
	var (
		req             []warehousedto.CreateClaimRequestRequest
		domainClaimReqs []warehouse.CreateClaimRequestItem
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	for _, value := range req {
		domainClaimReqs = append(domainClaimReqs, warehousedto.CreateClaimRequestDTOToDomain(value))
	}

	if err := h.uc.CreateClaimRequest(c.Request.Context(), domainClaimReqs, userID); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	c.JSON(http.StatusOK, warehousedto.CreateClaimRequestResponseWrapper{
		Code:    http.StatusCreated,
		Message: "Create claim request successfully",
		Data:    gin.H{},
	})
}

// @Summary Modify claim request
// @Description API chỉnh sửa thông tin đăng kí nhận đồ
// @Security BearerAuth
// @Tags item warehouses
// @Accept json
// @Produce json
// @Param request body warehousedto.ModifyClaimRequestRequest true "Claim request modify payload"
// @Success 201 {object} warehousedto.ModifyClaimRequestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/item-warehouses/claim-request [patch]
func (h *WarehouseHandler) ModifyClaimRequest(c *gin.Context) {
	var (
		req                      warehousedto.ModifyClaimRequestRequest
		domainModifyClaimRequest warehouse.ModifyClaimRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest),
		)
		return
	}

	domainModifyClaimRequest.ItemID = req.ItemID
	domainModifyClaimRequest.NewQuatity = req.NewQuantity

	if err := h.uc.ModifyClaimRequest(c.Request.Context(), domainModifyClaimRequest, userID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, warehousedto.ModifyClaimRequestResponseWrapper{
		Code:    http.StatusCreated,
		Message: "Modify claim request successfully",
		Data:    gin.H{},
	})
}

// @Summary Remove claim request
// @Description API hoàn tác đăng kí nhận đồ
// @Security BearerAuth
// @Tags item warehouses
// @Accept json
// @Produce json
// @Param itemID path int true "ID item"
// @Success 201 {object} warehousedto.RemoveClaimRequestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/item-warehouses/claim-request/{itemID} [delete]
func (h *WarehouseHandler) RemoveClaimRequest(c *gin.Context) {
	var (
		req warehousedto.RemoveClaimRequestRequest
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest),
		)
		return
	}

	if err := h.uc.RemoveClaimRequest(c.Request.Context(), req.ItemID, userID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, warehousedto.RemoveClaimRequestResponseWrapper{
		Code:    http.StatusCreated,
		Message: "Removed claim request successfully",
		Data:    gin.H{},
	})
}

// @Summary Remove all claim request
// @Description API xóa đăng kí nhận đồ
// @Security BearerAuth
// @Tags item warehouses
// @Accept json
// @Produce json
// @Success 201 {object} warehousedto.RemoveClaimRequestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/item-warehouses/claim-request [delete]
func (h *WarehouseHandler) RemoveAllClaimRequest(c *gin.Context) {
	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest),
		)
		return
	}

	if err := h.uc.RemoveAllClaimRequest(c.Request.Context(), userID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, warehousedto.RemoveClaimRequestResponseWrapper{
		Code:    http.StatusCreated,
		Message: "Removed all claim request successfully",
		Data:    gin.H{},
	})
}

// @Summary Get item warehouse by code
// @Description API lấy thông tin item warehouse bằng code
// @Security BearerAuth
// @Tags item warehouses
// @Accept json
// @Produce json
// @Param itemCode path int true "Code item warehouse"
// @Success 200 {object} warehousedto.GetItemWarehouseByCodeResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /item-warehouses/{itemCode} [get]
func (h *WarehouseHandler) GetItemByCode(c *gin.Context) {
	var (
		req                 warehousedto.GetItemWarehouseByCodeRequest
		domainItemWarehouse warehouse.ItemWareHouse
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.GetItemByCode(c.Request.Context(), &domainItemWarehouse, req.ItemCode); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var itemWarehouseDTORes warehousedto.ItemWareHouseDTO

	itemWarehouseDTORes = warehousedto.ItemWarehouseDomainToDTO(domainItemWarehouse)

	c.JSON(http.StatusOK, warehousedto.GetItemWarehouseByCodeResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched warehouses successfully",
		Data: warehousedto.GetItemWarehouseByCodeResponse{
			ItemWarehouse: itemWarehouseDTORes,
		},
	})
}

// @Summary Get warehouse by ID
// @Description API lấy thông tin warehouse theo ID
// @Security BearerAuth
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
		Message: "Fetched warehouse successfully",
		Data: warehousedto.GetWarehouseByIDResponse{
			Warehouse: warehouseDTORes,
		},
	})
}

// @Summary Update warehouse
// @Description API cập nhật warehouse
// @Security BearerAuth
// @Tags warehouses
// @Accept json
// @Produce json
// @Param warehouseID path int true "ID warehouse"
// @Param request body warehousedto.UpdateWarehouseRequest true "Update warehouse info"
// @Success 200 {object} warehousedto.UpdateWarehouseResponseWrapper "Updated warehouse successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /warehouses/{warehouseID} [patch]
func (h *WarehouseHandler) Update(c *gin.Context) {
	var (
		req             warehousedto.UpdateWarehouseRequest
		domainWarehouse warehouse.DetailWarehouse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	warehouseID, err := strconv.Atoi(c.Param("warehouseID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	domainWarehouse = warehousedto.UpdateWarehouseDTOToDomain(req)

	domainWarehouse.ID = uint(warehouseID)

	if err := h.uc.UpdateWarehouse(c.Request.Context(), domainWarehouse); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	c.JSON(http.StatusOK, warehousedto.UpdateWarehouseResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated warehouses successfully",
		Data:    gin.H{},
	})
}
