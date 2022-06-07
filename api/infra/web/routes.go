package web

import (
	"github.com/NicolasSales0101/ultividros-project/api/interface/adapters"
	"github.com/NicolasSales0101/ultividros-project/api/interface/controllers"
	"github.com/gofiber/fiber/v2"
)

// Função para executar o sentup de routes usando a biblioteca fiber
func Routes(app *fiber.App, c controllers.TemperedGlassController, ctx *fiber.Ctx) *fiber.App {
	TemperedGlassRoute := app.Group("/tempered-glasses")
	TemperedGlassRoute.Get("/id=:id", adapters.GetTemperedGlass(c.GetTemperedGlass, ctx))
	TemperedGlassRoute.Get("/", adapters.GetTemperedGlasses(c.GetTemperedGlasses, ctx))
	TemperedGlassRoute.Post("/", adapters.SaveTemperedGlasses(c.SaveTemperedGlasses, ctx))
	TemperedGlassRoute.Put("/", adapters.UpdateTemperedGlasses(c.UpdateTemperedGlasses, ctx))
	return app
}
