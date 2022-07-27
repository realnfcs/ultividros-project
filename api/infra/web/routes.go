package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/realnfcs/ultividros-project/api/infra/web/middlewares"
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
	TemperedGlassRoute.Post("/", middlewares.JWTAuth(), adapters.SaveTemperedGlasses(tempGlssControll.SaveTemperedGlasses, ctx))
	TemperedGlassRoute.Put("/", middlewares.JWTAuth(), adapters.UpdateTemperedGlasses(tempGlssControll.UpdateTemperedGlasses, ctx))
	TemperedGlassRoute.Patch("/", middlewares.JWTAuth(), adapters.PatchTemperedGlasses(tempGlssControll.PatchTemperedGlasses, ctx))
	TemperedGlassRoute.Delete("/", middlewares.JWTAuth(), adapters.DeleteTemperedGlass(tempGlssControll.DeleteTemperedGlass, ctx))

	// Common Glass Routes Section //
	comnGlssControll := c.CommonController

	CommonGlassRoute := app.Group("/common-glasses")
	CommonGlassRoute.Get("/id=:id", adapters.GetCommonGlass(comnGlssControll.GetCommonGlass, ctx))
	CommonGlassRoute.Get("/", adapters.GetCommonGlasses(comnGlssControll.GetCommonGlasses, ctx))
	CommonGlassRoute.Post("/", middlewares.JWTAuth(), adapters.SaveCommonGlasses(comnGlssControll.SaveCommonGlass, ctx))
	CommonGlassRoute.Patch("/", middlewares.JWTAuth(), adapters.PatchCommonGlasses(comnGlssControll.PatchCommonGlass, ctx))
	CommonGlassRoute.Delete("/", middlewares.JWTAuth(), adapters.DeleteCommonGlass(comnGlssControll.DeleteCommonGlass, ctx))

	// Part Routes Section //
	partControll := c.PartController

	PartRoute := app.Group("/parts")
	PartRoute.Get("/id=:id", adapters.GetPart(partControll.GetPart, ctx))
	PartRoute.Get("/", adapters.GetParts(partControll.GetParts, ctx))
	PartRoute.Post("/", middlewares.JWTAuth(), adapters.SavePart(partControll.SavePart, ctx))
	PartRoute.Patch("/", middlewares.JWTAuth(), adapters.PatchPart(partControll.PatchPart, ctx))
	PartRoute.Delete("/", middlewares.JWTAuth(), adapters.DeletePart(partControll.DeletePart, ctx))

	// User Routes Section //
	userControll := c.UserController

	UserRouter := app.Group("/users")
	UserRouter.Get("/id=:id", middlewares.JWTAuth(), adapters.GetUser(userControll.GetUser, ctx))
	UserRouter.Get("/", middlewares.JWTAuth(), adapters.GetUsers(userControll.GetUsers, ctx))
	UserRouter.Post("/", adapters.SaveUser(userControll.SaveUser, ctx))
	UserRouter.Patch("/", middlewares.JWTAuth(), adapters.PatchUser(userControll.PatchUser, ctx))
	UserRouter.Delete("/", middlewares.JWTAuth(), adapters.DeleteUser(userControll.DeleteUser, ctx))
	UserRouter.Post("/login", adapters.Login(userControll.Login, ctx))

	// Sale Routes Section //
	saleControll := c.SaleController

	SaleRouter := app.Group("/sales", middlewares.JWTAuth())
	SaleRouter.Get("/id=:id", adapters.GetSale(saleControll.GetSale, ctx))
	SaleRouter.Get("/", adapters.GetSales(saleControll.GetSales, ctx))
	SaleRouter.Post("/", adapters.SaveSale(saleControll.SaveSale, ctx))
	SaleRouter.Patch("/", adapters.PatchSale(saleControll.PatchSale, ctx))
	SaleRouter.Patch("/close", adapters.CloseSale(saleControll.CloseSale, ctx))
	SaleRouter.Delete("/", adapters.DeleteSale(saleControll.DeleteSale, ctx))

	return app
}
