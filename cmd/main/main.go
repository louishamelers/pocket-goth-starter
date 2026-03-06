package main

import (
	"log"
	"os"
	"pocket-goth-starter/internal/web/auth"
	middleware "pocket-goth-starter/internal/web/middleware"
	"pocket-goth-starter/internal/web/pages"
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

		initAuthRoutes(e)
		initAppRoutes(e)

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func initAuthRoutes(e *core.ServeEvent) {
	authGroup := e.Router.Group("/auth").BindFunc(middleware.UnAuthGuard)

	// Register
	authGroup.GET("/register", utils.RenderRoute(pages.RegisterPage))
	authGroup.POST("/register", auth.PostRegister)

	// Login
	authGroup.GET("/login", utils.RenderRoute(pages.LoginPage))
	authGroup.POST("/login", auth.PostLogin)

	// Logout (not part of authgroup)
	e.Router.POST("/auth/logout", auth.PostLogout)
}

func initAppRoutes(e *core.ServeEvent) {
	appGroup := e.Router.Group("/app").BindFunc(middleware.LoadAuthContext, middleware.AuthGuard)

	// Dashboard
	appGroup.GET("/dashboard", utils.RenderRoute(pages.DashboardPage))
}
