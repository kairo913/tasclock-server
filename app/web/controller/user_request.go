package controller

type UpdateUserRequest struct {
	Lastname  string `json:"lastname" validate:"max=20"`
	Firstname string `json:"firstname" validate:"max=20"`
}

type UpdateUserEmailRequest struct {
	Email string `json:"email" validate:"email"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" validate:"min=8,max=20"`
}
