package authdto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,password_strong" example:"Abc12345"`
	Device   string `json:"device" validate:"required,oneof=mobile web" example:"web"`
}
