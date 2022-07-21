package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realnfcs/ultividros-project/api/interface/adapters"
	"github.com/realnfcs/ultividros-project/api/interface/controllers"
)

// Função para executar o sentup de routes usando a biblioteca fiber
func Routes(app *fiber.App, c *controllers.Controllers, ctx *fiber.Ctx) *fiber.App {
	// Tempered Glass Routes Section //
	tempGlssControll := c.TemperedController

	TemperedGlassRoute := app.Group("/tempered-glasses")
	TemperedGlassRoute.Get("/id=:id", adapters.GetTemperedGlass(tempGlssControll.GetTemperedGlass, ctx))
	TemperedGlassRoute.Get("/", adapters.GetTemperedGlasses(tempGlssControll.GetTemperedGlasses, ctx))
	TemperedGlassRoute.Post("/", adapters.SaveTemperedGlasses(tempGlssControll.SaveTemperedGlasses, ctx))
	TemperedGlassRoute.Put("/", adapters.UpdateTemperedGlasses(tempGlssControll.UpdateTemperedGlasses, ctx))
	TemperedGlassRoute.Patch("/", adapters.PatchTemperedGlasses(tempGlssControll.PatchTemperedGlasses, ctx))
	TemperedGlassRoute.Delete("/", adapters.DeleteTemperedGlass(tempGlssControll.DeleteTemperedGlass, ctx))

	// Common Glass Routes Section //
	comnGlssControll := c.CommonController

	CommonGlassRoute := app.Group("/common-glasses")
	CommonGlassRoute.Get("/id=:id", adapters.GetCommonGlass(comnGlssControll.GetCommonGlass, ctx))
	CommonGlassRoute.Get("/", adapters.GetCommonGlasses(comnGlssControll.GetCommonGlasses, ctx))
	CommonGlassRoute.Post("/", adapters.SaveCommonGlasses(comnGlssControll.SaveCommonGlass, ctx))
	CommonGlassRoute.Patch("/", adapters.PatchCommonGlasses(comnGlssControll.PatchCommonGlass, ctx))
	CommonGlassRoute.Delete("/", adapters.DeleteCommonGlass(comnGlssControll.DeleteCommonGlass, ctx))

	// Part Routes Section //
	partControll := c.PartController

	PartRoute := app.Group("/parts")
	PartRoute.Get("/id=:id", adapters.GetPart(partControll.GetPart, ctx))
	PartRoute.Get("/", adapters.GetParts(partControll.GetParts, ctx))
	PartRoute.Post("/", adapters.SavePart(partControll.SavePart, ctx))
	PartRoute.Patch("/", adapters.PatchPart(partControll.PatchPart, ctx))
	PartRoute.Delete("/", adapters.DeletePart(partControll.DeletePart, ctx))

	// User Routes Section //
	userControll := c.UserController

	UserRouter := app.Group("/users")
	UserRouter.Get("/id=:id", adapters.GetUser(userControll.GetUser, ctx))
	UserRouter.Get("/", adapters.GetUsers(userControll.GetUsers, ctx))
	UserRouter.Post("/", adapters.SaveUser(userControll.SaveUser, ctx))
	UserRouter.Patch("/", adapters.PatchUser(userControll.PatchUser, ctx))
	UserRouter.Delete("/", adapters.DeleteUser(userControll.DeleteUser, ctx))

	// Sale Routes Section //
	saleControll := c.SaleController

	SaleRouter := app.Group("/sales")
	SaleRouter.Get("/id=:id", adapters.GetSale(saleControll.GetSale, ctx))
	SaleRouter.Get("/", adapters.GetSales(saleControll.GetSales, ctx))
	SaleRouter.Post("/", adapters.SaveSale(saleControll.SaveSale, ctx))
	SaleRouter.Patch("/", adapters.PatchSale(saleControll.PatchSale, ctx))

	return app
}
