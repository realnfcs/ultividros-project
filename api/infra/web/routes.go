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
	/*
		tempGlssControll := c.TemperedController

		TemperedGlassRoute := app.Group("/tempered-glasses")
		TemperedGlassRoute.Get("/id=:id", adapters.GetTemperedGlass(tempGlssControll.GetTemperedGlass, ctx))
		TemperedGlassRoute.Get("/", adapters.GetTemperedGlasses(tempGlssControll.GetTemperedGlasses, ctx))
		TemperedGlassRoute.Post("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.SaveTemperedGlasses(tempGlssControll.SaveTemperedGlasses, ctx))
		TemperedGlassRoute.Put("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.UpdateTemperedGlasses(tempGlssControll.UpdateTemperedGlasses, ctx))
		TemperedGlassRoute.Patch("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.PatchTemperedGlasses(tempGlssControll.PatchTemperedGlasses, ctx))
		TemperedGlassRoute.Delete("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.DeleteTemperedGlass(tempGlssControll.DeleteTemperedGlass, ctx))

		// Common Glass Routes Section //
		comnGlssControll := c.CommonController

		CommonGlassRoute := app.Group("/common-glasses")
		CommonGlassRoute.Get("/id=:id", adapters.GetCommonGlass(comnGlssControll.GetCommonGlass, ctx))
		CommonGlassRoute.Get("/", adapters.GetCommonGlasses(comnGlssControll.GetCommonGlasses, ctx))
		CommonGlassRoute.Post("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.SaveCommonGlasses(comnGlssControll.SaveCommonGlass, ctx))
		CommonGlassRoute.Patch("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.PatchCommonGlasses(comnGlssControll.PatchCommonGlass, ctx))
		CommonGlassRoute.Delete("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.DeleteCommonGlass(comnGlssControll.DeleteCommonGlass, ctx))

		// Part Routes Section //
		partControll := c.PartController

		PartRoute := app.Group("/parts")
		PartRoute.Get("/id=:id", adapters.GetPart(partControll.GetPart, ctx))
		PartRoute.Get("/", adapters.GetParts(partControll.GetParts, ctx))
		PartRoute.Post("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.SavePart(partControll.SavePart, ctx))
		PartRoute.Patch("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.PatchPart(partControll.PatchPart, ctx))
		PartRoute.Delete("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.DeletePart(partControll.DeletePart, ctx))

		// User Routes Section //
		userControll := c.UserController

		UserRouter := app.Group("/users")
		UserRouter.Get("/id=:id", middlewares.JWTAuth(), adapters.GetUser(userControll.GetUser, ctx))
		UserRouter.Get("/", middlewares.JWTAuth(), adapters.GetUsers(userControll.GetUsers, ctx))
		UserRouter.Post("/", adapters.SaveUser(userControll.SaveUser, ctx))
		UserRouter.Patch("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.PatchUser(userControll.PatchUser, ctx))
		UserRouter.Delete("/", middlewares.JWTAuth(), middlewares.JWTdata(), adapters.DeleteUser(userControll.DeleteUser, ctx))
		UserRouter.Post("/login", adapters.Login(userControll.Login, ctx), middlewares.CreateCookie())

		// Sale Routes Section //
		saleControll := c.SaleController

		SaleRouter := app.Group("/sales", middlewares.JWTAuth(), middlewares.JWTdata())
		SaleRouter.Get("/id=:id", adapters.GetSale(saleControll.GetSale, ctx))
		SaleRouter.Get("/", adapters.GetSales(saleControll.GetSales, ctx))
		SaleRouter.Post("/", adapters.SaveSale(saleControll.SaveSale, ctx))
		SaleRouter.Patch("/", adapters.PatchSale(saleControll.PatchSale, ctx))
		SaleRouter.Patch("/close", adapters.CloseSale(saleControll.CloseSale, ctx))
		SaleRouter.Delete("/", adapters.DeleteSale(saleControll.DeleteSale, ctx))
	*/

	tempGlssControll := c.TemperedController

	TemperedGlassRoute := app.Group("/tempered-glasses")
	TemperedGlassRoute.Get("/id=:id", adapters.GetTemperedGlass(tempGlssControll.GetTemperedGlass, ctx))
	TemperedGlassRoute.Get("/", adapters.GetTemperedGlasses(tempGlssControll.GetTemperedGlasses, ctx))
	TemperedGlassRoute.Post("/", middlewares.JWTAuthCookie(), adapters.SaveTemperedGlasses(tempGlssControll.SaveTemperedGlasses, ctx))
	TemperedGlassRoute.Put("/", middlewares.JWTAuthCookie(), adapters.UpdateTemperedGlasses(tempGlssControll.UpdateTemperedGlasses, ctx))
	TemperedGlassRoute.Patch("/", middlewares.JWTAuthCookie(), adapters.PatchTemperedGlasses(tempGlssControll.PatchTemperedGlasses, ctx))
	TemperedGlassRoute.Delete("/", middlewares.JWTAuthCookie(), adapters.DeleteTemperedGlass(tempGlssControll.DeleteTemperedGlass, ctx))

	// Common Glass Routes Section //
	comnGlssControll := c.CommonController

	CommonGlassRoute := app.Group("/common-glasses")
	CommonGlassRoute.Get("/id=:id", adapters.GetCommonGlass(comnGlssControll.GetCommonGlass, ctx))
	CommonGlassRoute.Get("/", adapters.GetCommonGlasses(comnGlssControll.GetCommonGlasses, ctx))
	CommonGlassRoute.Post("/", middlewares.JWTAuthCookie(), adapters.SaveCommonGlasses(comnGlssControll.SaveCommonGlass, ctx))
	CommonGlassRoute.Patch("/", middlewares.JWTAuthCookie(), adapters.PatchCommonGlasses(comnGlssControll.PatchCommonGlass, ctx))
	CommonGlassRoute.Delete("/", middlewares.JWTAuthCookie(), adapters.DeleteCommonGlass(comnGlssControll.DeleteCommonGlass, ctx))

	// Part Routes Section //
	partControll := c.PartController

	PartRoute := app.Group("/parts")
	PartRoute.Get("/id=:id", adapters.GetPart(partControll.GetPart, ctx))
	PartRoute.Get("/", adapters.GetParts(partControll.GetParts, ctx))
	PartRoute.Post("/", middlewares.JWTAuthCookie(), adapters.SavePart(partControll.SavePart, ctx))
	PartRoute.Patch("/", middlewares.JWTAuthCookie(), adapters.PatchPart(partControll.PatchPart, ctx))
	PartRoute.Delete("/", middlewares.JWTAuthCookie(), adapters.DeletePart(partControll.DeletePart, ctx))

	// User Routes Section //
	userControll := c.UserController

	UserRouter := app.Group("/users")
	UserRouter.Get("/id=:id", middlewares.JWTAuthCookie(), adapters.GetUser(userControll.GetUser, ctx))
	UserRouter.Get("/", middlewares.JWTAuthCookie(), adapters.GetUsers(userControll.GetUsers, ctx))
	UserRouter.Post("/", adapters.SaveUser(userControll.SaveUser, ctx))
	UserRouter.Patch("/", middlewares.JWTAuthCookie(), adapters.PatchUser(userControll.PatchUser, ctx))
	UserRouter.Delete("/", middlewares.JWTAuthCookie(), adapters.DeleteUser(userControll.DeleteUser, ctx))
	UserRouter.Post("/login", adapters.Login(userControll.Login, ctx), middlewares.CreateCookie())

	// Sale Routes Section //
	saleControll := c.SaleController

	SaleRouter := app.Group("/sales", middlewares.JWTAuthCookie())
	SaleRouter.Get("/id=:id", adapters.GetSale(saleControll.GetSale, ctx))
	SaleRouter.Get("/", adapters.GetSales(saleControll.GetSales, ctx))
	SaleRouter.Post("/", adapters.SaveSale(saleControll.SaveSale, ctx))
	SaleRouter.Patch("/", adapters.PatchSale(saleControll.PatchSale, ctx))
	SaleRouter.Patch("/close", adapters.CloseSale(saleControll.CloseSale, ctx))
	SaleRouter.Delete("/", adapters.DeleteSale(saleControll.DeleteSale, ctx))

	return app
}
