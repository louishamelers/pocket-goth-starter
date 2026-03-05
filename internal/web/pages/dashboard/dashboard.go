package dashboard

import (
	"pocket-goth-starter/internal/web/middleware"

	"github.com/pocketbase/pocketbase/core"
)

func GetDashboardRoute(e *core.RequestEvent) error {

	authRecord := e.Get(middleware.ContextAuthRecordKey)
	if authRecord == nil {
		e.Redirect(302, "/auth/register")
	}

	component := DashboardPage()
	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(e.Request.Context(), e.Response)
}
