package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	appEnv "pilotoAPI/internal/app/adapters"
	"pilotoAPI/internal/app/domain/repository"
	"pilotoAPI/internal/app/domain/routes"
	"pilotoAPI/internal/app/domain/routes/routedefinition"
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
	//services
	svcProduct := appEnv.NewProductService(repo)
	//handler
	handler := routes.NewHandler(svcProduct, log.Default())

	ec := echo.New()
	routedefinition.SetupHandlers(ec, handler)

	err = httpServer.Serve(ec, fmt.Sprintf(":%d", env.AppPort))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to start server: %s", err.Error()))
	}
}
