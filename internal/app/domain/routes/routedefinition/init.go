package routedefinition

import (
	"github.com/labstack/echo/v4"
	"pilotoAPI/internal/app/domain/routes"
)

func SetupHandlers(appEcho *echo.Echo, app *routes.Handler) {
	api := appEcho.Group("/api")
	routes.ProductAPIGroup(api, app)
}
