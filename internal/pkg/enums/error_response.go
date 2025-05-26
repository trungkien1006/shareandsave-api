package enums

type AppError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Name is require"`
	Error   string `json:"error" example:"INVALID_QUERY_PARAM"`
}

func NewAppError(code int, message string, err string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Error:   err,
	}
}

// Các lỗi phổ biến
var (
	ErrNotFound       = "ERR_NOT_FOUND"
	ErrUnauthorized   = "ERR_UNAUTHORIZED"
	ErrBadRequest     = "ERR_BAD_REQUEST"
	ErrInternalServer = "ERR_INTERNAL"
	ErrConflict       = "ERR_CONFLICT"
	ErrValidate       = "ERR_VALIDATE"
)
