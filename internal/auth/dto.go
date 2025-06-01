package auth

type SignInUserDto struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type SignUpUserDto struct {
	Name        string `json:"name" validate:"min=2,max=100"`
	Email       string `json:"email" validate:"email,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"min=8"`
	Role        string `json:"role" validate:"oneof=user admin"`
	ImageURL    string `json:"imageUrl,omitempty" validate:"omitempty,url"`
}

type SignInResponseDto struct {
	Token string `json:"token"`
}

type SignUpResponseDto struct {
	Token string `json:"token"`
}
