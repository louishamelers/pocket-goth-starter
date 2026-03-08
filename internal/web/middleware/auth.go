package middleware

import (
	"pocket-goth-starter/internal/web/routes"

	"github.com/pocketbase/pocketbase/core"
)

const AuthCookieName = "Auth"
const ContextAuthRecordKey = "authRecord"

func AuthGuard(e *core.RequestEvent) error {
	tokenCookie, err := e.Request.Cookie(AuthCookieName)
	if err != nil {
		// If no cookie is found, redirect to the login page
		e.Redirect(302, routes.LoginRoute)
		return nil
	}

	token := tokenCookie.Value
	if _, err := e.App.FindAuthRecordByToken(token, core.TokenTypeAuth); err != nil {
		// If the token is invalid, redirect to the login page
		e.Redirect(302, routes.LoginRoute)
		return nil
	}

	return e.Next()
}

func UnAuthGuard(e *core.RequestEvent) error {
	tokenCookie, err := e.Request.Cookie(AuthCookieName)

	if err != nil {
		return e.Next()
	}

	token := tokenCookie.Value
	if _, err := e.App.FindAuthRecordByToken(token, core.TokenTypeAuth); err != nil {
		return e.Next()
	}

	e.Redirect(302, routes.DashboardRoute)
	return nil
}

func LoadAuthContext(e *core.RequestEvent) error {
	tokenCookie, err := e.Request.Cookie(AuthCookieName)
	if err != nil {
		return e.Next()
	}

	token := tokenCookie.Value
	record, err := e.App.FindAuthRecordByToken(token, core.TokenTypeAuth)
	if err != nil {
		return e.Next()
	}

	e.Set(ContextAuthRecordKey, record)

	return e.Next()
}
