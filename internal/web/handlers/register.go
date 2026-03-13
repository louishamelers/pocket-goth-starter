package handlers

import (
	"pocket-goth-starter/internal/web/ui/pages"
	"pocket-goth-starter/internal/web/ui/utils"

	"github.com/pocketbase/pocketbase/core"
)

func HandleRegister() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		errorMessage := e.Request.URL.Query().Get("error")

		return utils.Render(e, (pages.RegisterPage(errorMessage)))
	}
}
