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

func PostRegisterRoute(e *core.RequestEvent) error {
	form := getRegisterFormValue(e)
	// TODO: validate form

	err := registerUser(e, form.email, form.password, form.passwordRepeat)

	if err != nil {
		fmt.Println(err)
	}

	return e.Redirect(302, "/app/dashboard")
}

func PostLogoutRoute(e *core.RequestEvent) error {
	fmt.Println("Hello")
	removeAuthToken(e)
	return e.Redirect(302, "/auth/login")
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
