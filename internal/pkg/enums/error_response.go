package enums

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code       string `json:"code" example:"INVALID_QUERY_PARAM"`
	Message    string `json:"message" example:"Tham số truy vấn không hợp lệ"`
	StatusCode int    `json:"status_code" example:"400"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func NewAppError(code string, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

// Các lỗi phổ biến
var (
	ErrNotFound       = NewAppError("ERR_NOT_FOUND", "Resource not found", http.StatusNotFound)
	ErrUnauthorized   = NewAppError("ERR_UNAUTHORIZED", "Unauthorized access", http.StatusUnauthorized)
	ErrBadRequest     = NewAppError("ERR_BAD_REQUEST", "Bad request", http.StatusBadRequest)
	ErrInternalServer = NewAppError("ERR_INTERNAL", "Internal server error", http.StatusInternalServerError)
	ErrConflict       = NewAppError("ERR_CONFLICT", "Conflict resource", http.StatusConflict)
)
