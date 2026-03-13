package auth

import "github.com/pocketbase/pocketbase/core"

type RegisterFormValue struct {
	Email          string
	Password       string
	PasswordRepeat string
}

type LoginFormValue struct {
	Email    string
	Password string
}

func GetLoginFormValue(e *core.RequestEvent) LoginFormValue {
	return LoginFormValue{
		Email:    e.Request.FormValue("email"),
		Password: e.Request.FormValue("password"),
	}
}

func GetRegisterFormValue(e *core.RequestEvent) RegisterFormValue {
	return RegisterFormValue{
		Email:          e.Request.FormValue("email"),
		Password:       e.Request.FormValue("password"),
		PasswordRepeat: e.Request.FormValue("passwordRepeat"),
	}
}
