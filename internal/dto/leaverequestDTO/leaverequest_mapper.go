package leaverequestdto

import leaverequests "final_project/internal/domain/leave_requests"

// Domain to DTO
func LeaveRequestDomainToDTO(domain leaverequests.LeaveRequest) LeaveRequestDTO {
	return LeaveRequestDTO{
		ID:        domain.ID,
		UserID:    domain.UserID,
		UserName:  domain.UserName,
		LeaveType: int(domain.LeaveType),
		Reason:    domain.Reason,
		StartDate: domain.StartDate,
		EndDate:   domain.EndDate,
		TotalDays: domain.TotalDays,

		CreatedAt: domain.CreatedAt,
	}
}
