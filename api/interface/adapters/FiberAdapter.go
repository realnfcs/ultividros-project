package adapters

import (
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/getcommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/savecommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/deletetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglasses"
	patchtemperedglass "github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/patchtempetedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/savetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/updatetemperedglass"
	"github.com/realnfcs/ultividros-project/api/interface/contracts"
)

// Tempered Glasses Section //

// Adaptador do fiber responsável pela obtenção de dados do vidro temperado no
// repositório de acordo com o parâmetro e pelo retorno dos mesmos para o cliente
func GetTemperedGlass[T contracts.FiberAdapterContract[T]](req func(gettemperedglass.Input) *gettemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		i := gettemperedglass.Input{ID: p}
		output := req(i)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por trazer ao cliente todos os dados dos vidros
// temperados no repositório
func GetTemperedGlasses[T contracts.FiberAdapterContract[T]](req func() *gettemperedglasses.Output, ctx T) func(T) error {
	return func(c T) error {
		output := req()
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por salvar o vidro temperado no repositório
func SaveTemperedGlasses[T contracts.FiberAdapterContract[T]](req func(savetemperedglass.Input) *savetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := savetemperedglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos de um vidro temperado no repositório
func UpdateTemperedGlasses[T contracts.FiberAdapterContract[T]](req func(updatetemperedglass.Input) *updatetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := updatetemperedglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos alterados de um vidro temperado no repositório
func PatchTemperedGlasses[T contracts.FiberAdapterContract[T]](req func(patchtemperedglass.Input) *patchtemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := patchtemperedglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por deletar o vidro temperado no repositório
func DeleteTemperedGlass[T contracts.FiberAdapterContract[T]](req func(deletetemperedglass.Input) *deletetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := deletetemperedglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Common Glasses Section //

// Adaptador do fiber responsável pela obtenção de dados do vidro comum no
// repositório de acordo com o parâmetro e pelo retorno dos mesmos para o cliente
func GetCommonGlass[T contracts.FiberAdapterContract[T]](req func(getcommonglass.Input) *getcommonglass.Output, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		i := getcommonglass.Input{ID: p}
		output := req(i)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por salvar o vidro comum no repositório
func SaveCommonGlasses[T contracts.FiberAdapterContract[T]](req func(savecommonglass.Input) *savecommonglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := savecommonglass.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}
