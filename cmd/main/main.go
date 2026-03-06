package main

import (
	"log"
	"os"
	"pocket-goth-starter/internal/web/auth"
	"pocket-goth-starter/internal/web/middleware"
	"pocket-goth-starter/internal/web/pages"
	"pocket-goth-starter/internal/web/routes"
	"pocket-goth-starter/internal/web/utils"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		// Serve static files
		// TODO: remove /tmp/
		e.Router.GET("/{path...}", apis.Static(os.DirFS("./tmp/pb_public"), false))

		initRoutes(e)

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func initRoutes(e *core.ServeEvent) {
	unAuthGroup := e.Router.Group("").BindFunc(middleware.UnAuthGuard)
	unAuthGroup.GET(routes.LoginRoute, utils.RenderRoute(pages.LoginPage))
	unAuthGroup.POST(routes.LoginRoute, auth.PostLogin)
	unAuthGroup.GET(routes.RegisterRoute, utils.RenderRoute(pages.RegisterPage))
	unAuthGroup.POST(routes.RegisterRoute, auth.PostRegister)

	e.Router.POST(routes.LogoutRoute, auth.PostLogout)

	authGroup := e.Router.Group("").BindFunc(middleware.LoadAuthContext, middleware.AuthGuard)
	authGroup.GET(routes.DashboardRoute, utils.RenderRoute(pages.DashboardPage))
}
