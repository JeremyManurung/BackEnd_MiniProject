package user

type RegisterUserInput struct {
	Nama       string `json:"nama" validate:"required"`
	Pekerjaan  string `json:"pekerjaan" validate:"required"`
	Email      string `json:"email" validate:"required, email"`
	Password   string `json:"password" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}