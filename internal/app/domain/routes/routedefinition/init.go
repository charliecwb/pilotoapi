package routedefinition

import (
	"github.com/gofiber/fiber/v2"
	"pilotoAPI/internal/app/domain/routes"
)

func SetupHandlers(appFiber fiber.Router, app *routes.Handler) {
	api := appFiber.Group("/api")
	routes.ProductAPIGroup(api, app)
}
