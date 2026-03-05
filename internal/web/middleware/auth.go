package middleware

import "github.com/pocketbase/pocketbase/core"

const AuthCookieName = "Auth"
const ContextAuthRecordKey = "authRecord"

func AuthGuard(e *core.RequestEvent) error {
	tokenCookie, err := e.Request.Cookie(AuthCookieName)
	if err != nil {
		// If no cookie is found, redirect to the login page
		e.Redirect(302, "/auth/login")
		return nil
	}

	token := tokenCookie.Value
	if _, err := e.App.FindAuthRecordByToken(token, core.TokenTypeAuth); err != nil {
		// If the token is invalid, redirect to the login page
		e.Redirect(302, "/auth/login")
		return nil
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

	e.Set(ContextAuthRecordKey, record)

	return e.Next()
}
