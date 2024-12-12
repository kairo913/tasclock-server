package controller

type SignUpRequest struct {
	Lastname  string `json:"lastname" validate:"max=50"`
	Firstname string `json:"firstname" validate:"required, max=50"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required, min=8, max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8, max=20"`
}
