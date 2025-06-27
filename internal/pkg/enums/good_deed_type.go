package enums

type GoodDeedType int

const (
	GoodDeedTypeAll          GoodDeedType = iota // 0 all
	GoodDeedTypeGiveOldItem                      // 1 tặng đồ cũ
	GoodDeedTypeGiveLoseItem                     // 2 trả đồ thất lạc
	GoodDeedTypeCampaign                         // 3 tham gia chiến dịch
)

func (s GoodDeedType) String() string {
	return [...]string{
		"ALL",
		"GIVE_AWAY_OLD_ITEM",
		"GIVE_LOSE_ITEM",
		"CAMPAGIN",
	}[s]
}
