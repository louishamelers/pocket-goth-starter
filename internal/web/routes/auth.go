package routes

import (
	"fmt"
	"pocket-goth-starter/internal/web/auth"

	"github.com/pocketbase/pocketbase/core"
)

func PostLogoutRoute(e *core.RequestEvent) error {
	fmt.Println("Hello")
	auth.RemoveAuthToken(e)
	return e.Redirect(302, "/auth/login")
}
