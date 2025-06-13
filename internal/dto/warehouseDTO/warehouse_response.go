package warehousedto

type FilterWarehouseResponseWrapper struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    FilterWarehouseResponse `json:"data"`
}

type FilterWarehouseResponse struct {
	Warehouses []WarehouseDTO `json:"warehouses"`
	TotalPage  int            `json:"totalPage"`
}

type GetWarehouseByIDResponseWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    GetWarehouseByIDResponse `json:"data"`
}

type GetWarehouseByIDResponse struct {
	Warehouse DetailWarehouseDTO `json:"warehouse"`
}

type UpdateWarehouseResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FilterItemWarehouseResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    FilterItemWarehouseResponse `json:"data"`
}

type FilterItemWarehouseResponse struct {
	ItemWarehouses []ItemWareHouseDTO `json:"itemWarehouses"`
	TotalPage      int                `json:"totalPage"`
}
