package auth

import (
	"net/url"
	"pocket-goth-starter/internal/web/routes"

	"github.com/pocketbase/pocketbase/core"
)

func PostLogin(e *core.RequestEvent) error {
	form := GetLoginFormValue(e)
	// TODO: validate form

	err := LoginUser(e, form.Email, form.Password)

	if err != nil {
		return e.Redirect(302, routes.LoginRoute+"?error=invalid_credentials")
	}

	return e.Redirect(302, routes.DashboardRoute)
}

func PostRegister(e *core.RequestEvent) error {
	form := GetRegisterFormValue(e)
	// TODO: validate form

	err := RegisterUser(e, form.Email, form.Password, form.PasswordRepeat)

	if err != nil {
		return e.Redirect(302, routes.RegisterRoute+"?error="+url.QueryEscape(err.Error()))
	}

	return e.Redirect(302, routes.DashboardRoute)
}

func PostLogout(e *core.RequestEvent) error {
	LogoutUser(e)
	return e.Redirect(302, routes.LoginRoute)
}
