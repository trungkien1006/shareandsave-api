package enums

type PostType int

const (
	PostTypeAll             PostType = iota // 0 all
	PostTypeGiveAwayOldItem                 // 1 tặng đồ cũ
	PostTypeFoundItem                       // 2 nhặt được đồ
	PostTypeSeekLoseItem                    // 3 tìm kiếm đồ
	PostTypeOther                           // 4 khác
)

func (s PostType) String() string {
	return [...]string{"ALL", "GIVE_AWAY_OLD_ITEM", "FOUND_ITEM", "SEEK_LOSE_ITEM", "OTHER"}[s]
}
