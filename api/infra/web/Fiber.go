// Pacote responsável pela a parte web da API e fica na
// camada Frameworks and Drivers da Clean Architecture
package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realnfcs/ultividros-project/api/infra/repository"
	"github.com/realnfcs/ultividros-project/api/interface/controllers"
)

// Função para iniciar o servidor Fiber
func Fiber() *fiber.App {
	app := fiber.New()
	repo, err := new(repository.TemperedGlassRepositoryMySql).Init()
	if err != nil {
		panic("Cannot inicialize the database")
	}

	controllers := controllers.TemperedGlassController{Repo: repo}
	Routes(app, controllers, new(fiber.Ctx))

	return app
}
