package routes

import (
	"github.com/labstack/echo/v4"
)

func ProductAPIGroup(appEcho *echo.Group, app *Handler) {
	api := appEcho.Group("/product")

	api.GET("/:id", app.fetchProduct)
	api.GET("/", app.fetchAllProducts)
	api.POST("/", app.createProduct)
	api.PUT("/:id", app.updateProduct)
	api.DELETE("/:id", app.deleteProduct)
}
