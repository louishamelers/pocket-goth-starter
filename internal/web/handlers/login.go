package handlers

import (
	"pocket-goth-starter/internal/web/ui/pages"
	"pocket-goth-starter/internal/web/ui/utils"

	"github.com/pocketbase/pocketbase/core"
)

func HandleLogin() func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		errorMessage := ""
		if e.Request.URL.Query().Get("error") != "" {
			errorMessage = "Invalid credentials"
		}

		return utils.Render(e, (pages.LoginPage(errorMessage)))
	}
}
