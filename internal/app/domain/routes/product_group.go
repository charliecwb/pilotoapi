package routes

import "github.com/gofiber/fiber/v2"

func ProductAPIGroup(appFiber fiber.Router, app *Handler) {
	api := appFiber.Group("/product")

	api.Get("/:id", app.fetchProduct)
	api.Get("/", app.fetchAllProducts)
	api.Post("/", app.createProduct)
	api.Put("/:id", app.updateProduct)
	api.Delete("/:id", app.deleteProduct)
}
