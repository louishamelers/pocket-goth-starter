package handlers

import (
	"pocket-goth-starter/internal/web/pages"
	"pocket-goth-starter/internal/web/utils"

	"github.com/pocketbase/pocketbase/core"
)

func HandleRegister() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		errorMessage := e.Request.URL.Query().Get("error")

		return utils.Render(e, (pages.RegisterPage(errorMessage)))
	}
}
