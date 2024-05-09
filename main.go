package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v "github.com/stelgkio/otoo_landing/view"
	cu "github.com/stelgkio/otoo_landing/view/component/customer/overview"
	ex "github.com/stelgkio/otoo_landing/view/component/extension/overview"
	p "github.com/stelgkio/otoo_landing/view/component/product/overview"
	d "github.com/stelgkio/otoo_landing/view/dashboard"
	t "github.com/stelgkio/otoo_landing/view/dashboard/template"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	e.Static("/css", "css")
	e.Static("/img", "img")
	e.Static("/js", "js")
	e.Static("/fonts", "fonts")

	e.GET("/", func(c echo.Context) error {
		return Render(c, v.IndexTemplate())
	})
	e.GET("/dashboard", func(c echo.Context) error {
		return Render(c, t.DeafultTemplate())
	})
	e.GET("/dashboard/default", func(c echo.Context) error {
		return Render(c, t.DeafultTemplate())
	})
	e.GET("/dashboard/analytics", func(c echo.Context) error {
		return Render(c, t.AnalyticsTemplate())
	})
	e.GET("/dashboard/wallet", func(c echo.Context) error {
		return Render(c, t.WalletTemplate())
	})
	e.GET("/default", func(c echo.Context) error {
		return Render(c, d.DeafultDashboard())
	})
	e.GET("/analytics", func(c echo.Context) error {
		return Render(c, d.AnalyticsDashboard())
	})
	e.GET("/wallet", func(c echo.Context) error {
		return Render(c, d.WalletDashboard())
	})
	e.GET("/product", func(c echo.Context) error {
		return Render(c, p.ProductOverview())
	})
	e.GET("/customer", func(c echo.Context) error {
		return Render(c, cu.CustomerOverView())
	})
	e.GET("/extension", func(c echo.Context) error {
		return Render(c, ex.ExtensionOverview())
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}
