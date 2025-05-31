package enums

type PostType int

const (
	PostTypeGiveAwayOldItem PostType = iota // 0 tặng đồ cũ
	PostTypeFoundItem                       // 1 nhặt được đồ
	PostTypeSeekLoseItem                    // 2 tìm kiếm đồ
	PostTypeOther                           // 3 khác
)

func (s PostType) String() string {
	return [...]string{"GIVE_AWAY_OLD_ITEM", "FOUND_ITEM", "SEEK_LOSE_ITEM", "OTHER"}[s]
}
