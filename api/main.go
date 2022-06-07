package main

import "github.com/realnfcs/ultividros-project/api/infra/web"

func main() {
	app := web.Fiber()

	app.Listen(":3000")
}
