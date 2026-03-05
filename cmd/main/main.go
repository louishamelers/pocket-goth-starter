package main

import (
	"log"
	"os"
	"pocket-goth-starter/internal/web/components"
	"pocket-goth-starter/internal/web/middleware"
	dashboardPage "pocket-goth-starter/internal/web/pages/dashboard"
	loginPage "pocket-goth-starter/internal/web/pages/login"
	registerPage "pocket-goth-starter/internal/web/pages/register"

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

		// Other routes
		e.Router.GET("/hello/{name}", helloRoute)

		authGroup := e.Router.Group("/auth")

		// Register
		authGroup.GET("/register", registerPage.GetRegisterRoute)
		authGroup.POST("/register", registerPage.PostRegisterRoute)

		// Login
		authGroup.GET("/login", loginPage.GetLoginRoute)
		authGroup.POST("/login", loginPage.PostLoginRoute)

		// Logout
		authGroup.POST("/logout", registerPage.PostLogoutRoute)

		appGroup := e.Router.Group("/app").BindFunc(func(e *core.RequestEvent) error {
			tokenCookie, err := e.Request.Cookie(middleware.AuthCookieName)
			if err != nil {
				return e.Next()
			}

			token := tokenCookie.Value
			record, err := e.App.FindAuthRecordByToken(token, core.TokenTypeAuth)
			if err != nil {
				return e.Next()
			}

			e.Set(middleware.ContextAuthRecordKey, record)

			return e.Next()
		})

		// Dashboard
		appGroup.GET("/dashboard", dashboardPage.GetDashboardRoute)

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func helloRoute(e *core.RequestEvent) error {
	name := e.Request.PathValue("name")

	component := components.Index(name)

	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")

	return component.Render(e.Request.Context(), e.Response)
}
