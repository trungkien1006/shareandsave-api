package enums

type ItemClassify int

const (
	ItemClassifyAll      ItemClassify = iota // 0 all
	ItemClassifyOlItem                       // 1 đồ cũ
	ItemClassifyLoseItem                     // 2 đồ thất lạc
)

func (s ItemClassify) String() string {
	return [...]string{"ALL", "OLD_ITEM", "LOSE_ITEM"}[s]
}
