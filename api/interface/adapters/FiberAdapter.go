package adapters

import (
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglasses"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/savetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/updatetemperedglass"
	"github.com/realnfcs/ultividros-project/api/interface/contracts"
)

// Adaptador do fiber responsável pela obtenção de dados do vidro temperado no
// repositório de acordo com o parâmetro e pelo retorno dos mesmos para o cliente
func GetTemperedGlass[T contracts.Adapters](req func(gettemperedglass.Input) *gettemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		i := gettemperedglass.Input{ID: p}
		obj := req(i)
		return c.JSON(obj)
	}
}

// Adaptador do fiber responsável por trazer ao cliente todos os dados dos vidros
// temperados no repositório
func GetTemperedGlasses[T contracts.Adapters](req func() *gettemperedglasses.Output, ctx T) func(T) error {
	return func(c T) error {
		obj := req()
		return c.JSON(obj)
	}
}

// Adaptador do fiber responsável por salvar o vidro temperado no repositório
func SaveTemperedGlasses[T contracts.Adapters](req func(savetemperedglass.Input) *savetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := savetemperedglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.JSON(output)
	}
}

func UpdateTemperedGlasses[T contracts.Adapters](req func(updatetemperedglass.Input) *updatetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := updatetemperedglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.JSON(output)
	}
}
