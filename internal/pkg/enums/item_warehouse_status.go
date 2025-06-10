package enums

type ItemWarehouseStatus int

const (
	ItemWarehouseStatusAll      ItemWarehouseStatus = iota // 0
	ItemWarehouseStatusInStock                             // 1 Đồ còn trong kho
	ItemWarehouseStatusOutStock                            // 2 Đồ đã xuất kho
)

func (s ItemWarehouseStatus) String() string {
	return [...]string{"ALL", "IN_STOCK", "OUT_STOCK"}[s]
}
