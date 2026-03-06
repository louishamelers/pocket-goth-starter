package auth

import (
	"fmt"
	"pocket-goth-starter/internal/web/routes"

	"github.com/pocketbase/pocketbase/core"
)

func PostLogin(e *core.RequestEvent) error {
	form := GetLoginFormValue(e)
	// TODO: validate form

	err := LoginUser(e, form.Email, form.Password)

	if err != nil {
		fmt.Println(err)
	}

	return e.Redirect(302, routes.DashboardRoute)
}

func PostRegister(e *core.RequestEvent) error {
	form := GetRegisterFormValue(e)
	// TODO: validate form

	err := RegisterUser(e, form.Email, form.Password, form.PasswordRepeat)

	if err != nil {
		fmt.Println(err)
	}

	return e.Redirect(302, routes.DashboardRoute)
}

func PostLogout(e *core.RequestEvent) error {
	LogoutUser(e)
	return e.Redirect(302, routes.LoginRoute)
}
