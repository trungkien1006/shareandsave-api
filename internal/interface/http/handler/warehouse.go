package handler

import "final_project/internal/application/warehouseapp"

type WarehouseHandler struct {
	uc *warehouseapp.UseCase
}

func NewWarehouseHandler(uc *warehouseapp.UseCase) *WarehouseHandler {
	return &WarehouseHandler{uc: uc}
}
