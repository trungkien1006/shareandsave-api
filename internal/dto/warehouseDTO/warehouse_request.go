package warehousedto

type GetWarehouseRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=createdAt" example:"createdAt quantity"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"`
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=senderName invoiceNum itemName classify sku description stockPlace"`
	SearchValue string `form:"searchValue"`
}

func (r *GetWarehouseRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

type GetWarehouseByIDRequest struct {
	WarehouseID uint `uri:"warehouseID" binding:"required"`
}

type UpdateWarehouseRequest struct {
	Description    string                `json:"description"`
	StockPlace     string                `json:"stockPlace"`
	ItemWarehouses []UpdateItemWarehouse `json:"itemWarehouses"`
}

type UpdateItemWarehouse struct {
	ID         uint   `json:"id" binding:"required"`
	Descripton string `json:"description" binding:"required"`
}
