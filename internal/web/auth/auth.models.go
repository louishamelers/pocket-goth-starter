package auth

type RegisterFormValue struct {
	Email          string
	Password       string
	PasswordRepeat string
}

type LoginFormValue struct {
	Email    string
	Password string
}
