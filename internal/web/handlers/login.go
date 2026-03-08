package handlers

import (
	"pocket-goth-starter/internal/web/pages"
	"pocket-goth-starter/internal/web/utils"

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
