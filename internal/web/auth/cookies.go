package auth

import (
	"net/http"
	"pocket-goth-starter/internal/web/middleware"
	"time"

	"github.com/pocketbase/pocketbase/core"
)

func RemoveAuthToken(e *core.RequestEvent) {
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

func SetAuthToken(e *core.RequestEvent, user *core.Record) error {
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
