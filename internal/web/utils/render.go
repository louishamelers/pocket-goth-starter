package utils

import (
	"github.com/a-h/templ"
	"github.com/pocketbase/pocketbase/core"
)

func Render(e *core.RequestEvent, component templ.Component) error {
	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	return component.Render(e.Request.Context(), e.Response)
}

func RenderRoute(componentFunc func() templ.Component) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		return Render(e, componentFunc())
	}
}
