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
