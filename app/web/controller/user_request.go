package controller

type SignUpRequest struct {
	Lastname  string `json:"lastname" validate:"max=20"`
	Firstname string `json:"firstname" validate:"required, max=20"`
	Email     string `json:"email" validate:"required, email"`
	Password  string `json:"password" validate:"required, min=6"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=6"`
}
