package register

import (
	"fmt"
	"net/http"
	"pocket-goth-starter/internal/web/middleware"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

type RegisterFormValue struct {
	email          string
	password       string
	passwordRepeat string
}

func getRegisterFormValue(e *core.RequestEvent) RegisterFormValue {
	return RegisterFormValue{
		email:          e.Request.FormValue("email"),
		password:       e.Request.FormValue("password"),
		passwordRepeat: e.Request.FormValue("passwordRepeat"),
	}
}

func GetRegisterRoute(e *core.RequestEvent) error {
	component := RegisterPage()
	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(e.Request.Context(), e.Response)
}

func PostRegisterRoute(e *core.RequestEvent) error {
	form := getRegisterFormValue(e)
	// TODO: validate form

	err := registerUser(e, form.email, form.password, form.passwordRepeat)

	if err != nil {
		fmt.Println(err)
	}

	// just show the registerPage
	component := RegisterPage()
	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(e.Request.Context(), e.Response)
}

func PostLogoutRoute(e *core.RequestEvent) error {
	removeAuthToken(e)
	return e.Redirect(302, "/auth/register")
}

func registerUser(e *core.RequestEvent, email string, password string, repeatPassword string) error {
	user, _ := e.App.FindAuthRecordByEmail("users", email)
	if user != nil {
		return fmt.Errorf("email already exists!")
	}

	// TODO: move this to validation
	if repeatPassword != password {
		return fmt.Errorf("passwords do not match!")
	}

	userCollection, err := e.App.FindCollectionByNameOrId("users")
	if err != nil {
		return err
	}

	newUser := core.NewRecord(userCollection)
	newUser.SetPassword(password)
	newUser.SetEmail(email)

	if err := e.App.Save(newUser); err != nil {
		return err
	}

	return setAuthToken(e, newUser)
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
