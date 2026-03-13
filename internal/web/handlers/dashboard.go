package handlers

import (
	"pocket-goth-starter/internal/web/ui/pages"
	"pocket-goth-starter/internal/web/ui/utils"

	"github.com/pocketbase/pocketbase/core"
)

func HandleDashbaord() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		accountRecord := e.Auth
		accountEmail := accountRecord.GetString("email")
		return utils.Render(e, (pages.DashboardPage(accountEmail)))
	}
}
