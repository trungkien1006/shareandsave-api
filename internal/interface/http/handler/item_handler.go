package handler

import (
	"final_project/internal/application/itemapp"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/item"
	itemdto "final_project/internal/dto/itemDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	uc *itemapp.UseCase
}

func NewItemHandler(uc *itemapp.UseCase) *ItemHandler {
	return &ItemHandler{uc: uc}
}

// @Summary Get items
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags items
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Param sort query string false "Sort column (vd: name)" example(name)
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   searchBy   query    string  false  "Trường lọc (vd: email, full_name)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} itemdto.GetItemResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /items [get]
func (h *ItemHandler) GetAllItem(c *gin.Context) {
	var req itemdto.GetAllItemRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	req.SetDefault()

	var items []item.Item
	domainReq := filter.FilterRequest{
		Page:        req.Page,
		Limit:       req.Limit,
		Sort:        req.Sort,
		Order:       req.Order,
		SearchBy:    req.SearchBy,
		SearchValue: req.SearchValue,
	}

	totalPage, err := h.uc.GetAllItem(c.Request.Context(), &items, domainReq)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	itemdtos := make([]itemdto.ItemDTO, 0)

	for _, i := range items {
		itemdtos = append(itemdtos, itemdto.ToItemDTO(i))
	}

	c.JSON(http.StatusOK, itemdto.GetItemResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched items successfully",
		Data: itemdto.GetItemResponse{
			Items:     itemdtos,
			TotalPage: totalPage,
		},
	})
}

// @Summary Get item by ID
// @Description API lấy thông tin item theo ID
// @Tags items
// @Accept json
// @Produce json
// @Param itemID path int true "ID item"
// @Success 200 {object} itemdto.GetItemByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /items/{itemID} [get]
func (h *ItemHandler) GetItemByID(c *gin.Context) {
	var req itemdto.GetItemByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}
	var itm item.Item
	if err := h.uc.GetItemByID(c.Request.Context(), &itm, req.ItemID); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}
	c.JSON(http.StatusOK, itemdto.GetItemByIDResponseWrapper{
		Code:    http.StatusOK,
		Message: "Item fetched successfully",
		Data: itemdto.GetItemByIDResponse{
			Item: itemdto.ToItemDTO(itm),
		},
	})
}

// @Summary Create new item
// @Description API thêm item mới
// @Tags items
// @Accept json
// @Produce json
// @Param request body itemdto.CreateItemRequest true "Create item info"
// @Success 200 {object} itemdto.CreateItemResponseWrapper "Created item successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /items [post]
func (h *ItemHandler) CreateItem(c *gin.Context) {
	var req itemdto.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}
	itm := item.Item{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	}
	if err := h.uc.CreateItem(c.Request.Context(), &itm); err != nil {
		c.JSON(http.StatusInternalServerError, enums.NewAppError(http.StatusInternalServerError, err.Error(), enums.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, itemdto.CreateItemResponseWrapper{
		Code:    http.StatusOK,
		Message: "Item created successfully",
		Data: itemdto.CreateItemResponse{
			Item: itemdto.ToItemDTO(itm),
		},
	})
}

// @Summary Update item
// @Description API cập nhật item
// @Tags items
// @Accept json
// @Produce json
// @Param request body itemdto.UpdateItemRequest true "Update item info"
// @Success 200 {object} itemdto.UpdateItemResponseWrapper "Updated item successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /items [put]
func (h *ItemHandler) UpdateItem(c *gin.Context) {
	var req itemdto.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}
	itm := &item.Item{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
	}
	if err := h.uc.UpdateItem(c.Request.Context(), itm); err != nil {
		c.JSON(http.StatusInternalServerError, enums.NewAppError(http.StatusInternalServerError, err.Error(), enums.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, itemdto.UpdateItemResponseWrapper{
		Code:    http.StatusOK,
		Message: "Item updated successfully",
		Data: itemdto.UpdateItemResponse{
			Item: itemdto.ToItemDTO(*itm),
		},
	})
}

// @Summary Delete item
// @Description API xóa item theo ID
// @Tags items
// @Accept json
// @Produce json
// @Param itemID path int true "ID item"
// @Success 200 {object} itemdto.DeleteItemResponseWrapper "Deleted item successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /items/{itemID} [delete]
func (h *ItemHandler) DeleteItem(c *gin.Context) {
	var req itemdto.DeleteItemRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}
	if err := h.uc.DeleteItem(c.Request.Context(), req.ItemID); err != nil {
		c.JSON(http.StatusInternalServerError, enums.NewAppError(http.StatusInternalServerError, err.Error(), enums.ErrInternal))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Item deleted successfully",
	})
}
