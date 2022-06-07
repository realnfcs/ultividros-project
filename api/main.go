package main

import "github.com/NicolasSales0101/ultividros-project/api/infra/web"

func main() {
	app := web.Fiber()

	app.Listen(":3000")
}
