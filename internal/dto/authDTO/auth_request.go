package authdto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"superadmin@example.com"`
	Password string `json:"password" validate:"required,password_strong" example:"Admin1234"`
	// Password string `json:"password" validate:"required" example:"Abc12345"`

	Device string `json:"device" validate:"required,oneof=mobile web" example:"web"`
}

type GetAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type SignUpRequest struct {
	VerifyToken string `json:"verifyToken" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phone_vn"`
	FullName    string `json:"fullName" validate:"required,min=2,max=50"`
	Password    string `json:"password" validate:"required,password_strong"`
	RePassword  string `json:"rePassword" validate:"required,eqfield=Password"`
}

type VerifySignUpRequest struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phone_vn"`
	FullName    string `json:"fullName" validate:"required,min=2,max=50"`
	Password    string `json:"password" validate:"required,password_strong"`
	RePassword  string `json:"rePassword" validate:"required,eqfield=Password"`
}

type ResetPasswordRequest struct {
	VerifyToken     string `json:"verifyToken" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	CurrentPassword string `json:"currentPassword"`
	Password        string `json:"password" validate:"required,password_strong"`
	RePassword      string `json:"rePassword" validate:"required,eqfield=Password"`
}

type SendOTPRequest struct {
	Email   string `json:"email" validate:"required,email"`
	Purpose string `json:"purpose" validate:"required,oneof=activeAccount resetPassword"`
}

type VerifyOTPRequest struct {
	Email   string `json:"email" validate:"required,email"`
	Purpose string `json:"purpose" validate:"required,oneof=activeAccount resetPassword"`
	OTP     string `json:"otp" validate:"required"`
}
