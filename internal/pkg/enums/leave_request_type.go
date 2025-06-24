package enums

type LeaveRequestType int

const (
	LeaveRequestTypeAll          LeaveRequestType = iota // 0 Tất cả
	LeaveRequestTypeAnnual                               // 1 Nghỉ phép năm
	LeaveRequestTypeSick                                 // 2 Nghỉ bệnh
	LeaveRequestTypeMaternity                            // 3 Nghỉ thai sản
	LeaveRequestTypePaternity                            // 4 Nghỉ thai sản (cha)
	LeaveRequestTypeMarriage                             // 5 Nghỉ cưới
	LeaveRequestTypeBereavement                          // 6 Nghỉ tang
	LeaveRequestTypeUnpaid                               // 7 Nghỉ không lương
	LeaveRequestTypePersonal                             // 8 Nghỉ cá nhân
	LeaveRequestTypeSpecial                              // 9 Nghỉ đặc biệt
	LeaveRequestTypeParental                             // 10 Nghỉ chăm con
	LeaveRequestTypeCompensatory                         // 11 Nghỉ bù
	LeaveRequestTypeBusiness                             // 12 Nghỉ công tác
	LeaveRequestTypeRecuperation                         // 13 Nghỉ dưỡng sức
)

func (s LeaveRequestType) String() string {
	return [...]string{
		"ALL",
		"ANNUAL_LEAVE",
		"SICK_LEAVE",
		"MATERNITY_LEAVE",
		"PATERNITY_LEAVE",
		"MARRIAGE_LEAVE",
		"BEREAVEMENT_LEAVE",
		"UNPAID_LEAVE",
		"PERSONAL_LEAVE",
		"SPECIAL_LEAVE",
		"PARENTAL_LEAVE",
		"COMPENSATORY_LEAVE",
		"BUSINESS_LEAVE",
		"RECUPERATION_LEAVE",
	}[s]
}
