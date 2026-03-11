package middleware

import (
	"pocket-goth-starter/internal/web/routes"

	"github.com/pocketbase/pocketbase/core"
)

const (
	AuthCookieName       = "Auth"
	ContextAuthRecordKey = "authRecord"
)

func AuthGuard(e *core.RequestEvent) error {
	if e.Auth == nil {
		return e.Redirect(302, routes.LoginRoute)
	}
	return e.Next()
}

func UnAuthGuard(e *core.RequestEvent) error {
	if e.Auth != nil {
		return e.Redirect(302, routes.DashboardRoute)
	}
	return e.Next()
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

	e.Auth = record

	return e.Next()
}
