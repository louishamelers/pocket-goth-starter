package login

import (
	"fmt"
	"net/http"
	"pocket-goth-starter/internal/web/middleware"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

type LoginFormValue struct {
	email    string
	password string
}

func getLoginFormValue(e *core.RequestEvent) LoginFormValue {
	return LoginFormValue{
		email:    e.Request.FormValue("email"),
		password: e.Request.FormValue("password"),
	}
}

func PostLoginRoute(e *core.RequestEvent) error {
	form := getLoginFormValue(e)
	// TODO: validate form

	err := loginUser(e, form.email, form.password)

	if err != nil {
		fmt.Println(err)
	}

	return e.Redirect(302, "/app/dashboard")
}

func loginUser(e *core.RequestEvent, email string, password string) error {
	user, err := e.App.FindAuthRecordByEmail("users", email)
	if err != nil {
		return fmt.Errorf("Login failed")
	}

	valid := user.ValidatePassword(password)

	if !valid {
		return fmt.Errorf("Login failed")
	}

	return setAuthToken(e, user)
}

func setAuthToken(e *core.RequestEvent, user *core.Record) error {
	token, err := user.NewAuthToken()
	if err != nil {
		return err
	}
	e.SetCookie(&http.Cookie{
		Name:     middleware.AuthCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set to true except in dev mode
		SameSite: http.SameSiteLaxMode,
		MaxAge:   int(time.Hour * 24 * 7 / time.Second), // 7 days
	})
	return nil
}

func removeAuthToken(e *core.RequestEvent) {
	e.SetCookie(&http.Cookie{
		Name:     middleware.AuthCookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // TODO: set to true except in dev mode
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
}
