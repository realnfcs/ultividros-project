// Pacote responsável pela a parte web da API e fica na
// camada Frameworks and Drivers da Clean Architecture
package web

import (
	"github.com/NicolasSales0101/ultividros-project/api/infra/repository"
	"github.com/NicolasSales0101/ultividros-project/api/interface/controllers"
	"github.com/gofiber/fiber/v2"
)

// Função para iniciar o servidor Fiber
func Fiber() *fiber.App {
	app := fiber.New()

	repo := new(repository.TemperedGlassRepositoryMemory).Init()

	controllers := controllers.TemperedGlassController{Repo: repo}
	//fiberAdapter := new(adapters.Adapter[*fiber.Ctx]).Init()

	Routes(app, controllers, new(fiber.Ctx))

	return app
}
