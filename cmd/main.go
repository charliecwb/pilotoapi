package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	appEnv "pilotoAPI/internal/app/adapters"
	"pilotoAPI/internal/app/domain/repository"
	"pilotoAPI/internal/app/domain/routes"
	"pilotoAPI/internal/app/domain/routes/routedefinition"
	"pilotoAPI/internal/app/services"
	"pilotoAPI/internal/pkg/server"
)

func main() {
	err := appEnv.NewEnv(appEnv.DEFULT_ENV)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create env: %s", err.Error()))
	}
	env := appEnv.GetEnv()

	httpServer := server.NewHTTPServer()

	db, err := server.NewDataBase(env.Env)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to connect in db: %s", err.Error()))
	}
	//repo
	repo := repository.NewProductRepository(db)
	//service
	svcProduct := services.NewProductService(repo)
	//handler
	handler := routes.NewHandler(svcProduct, log.Default())
	app := fiber.New()
	routedefinition.SetupHandlers(app, handler)

	err = httpServer.Serve(app, fmt.Sprintf(":%d", env.AppPort))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to start server: %s", err.Error()))
	}
}
