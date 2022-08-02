package adapters

import (
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/deletecommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/getcommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/getcommonglasses"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/patchcommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/commonglasses/savecommonglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/deletepart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/getpart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/getparts"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/patchpart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/parts/savepart"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/closesale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/deletesale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/getsale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/getsales"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/patchsale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/sales/savesale"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/deletetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglasses"
	patchtemperedglass "github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/patchtempetedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/savetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/updatetemperedglass"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/deleteuser"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/getuser"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/getusers"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/login"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/patchuser"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/saveuser"
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
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos de um vidro temperado no repositório
func UpdateTemperedGlasses[T contracts.FiberAdapterContract[T]](req func(updatetemperedglass.Input) *updatetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := updatetemperedglass.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos alterados de um vidro temperado no repositório
func PatchTemperedGlasses[T contracts.FiberAdapterContract[T]](req func(patchtemperedglass.Input) *patchtemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := patchtemperedglass.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por deletar o vidro temperado no repositório
func DeleteTemperedGlass[T contracts.FiberAdapterContract[T]](req func(deletetemperedglass.Input) *deletetemperedglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := deletetemperedglass.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

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

// Adaptador do fiber responsável por trazer ao cliente todos os dados dos vidros
// comuns no repositório
func GetCommonGlasses[T contracts.FiberAdapterContract[T]](req func() *getcommonglasses.Output, ctx T) func(T) error {
	return func(c T) error {
		output := req()
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por salvar o vidro comum no repositório
func SaveCommonGlasses[T contracts.FiberAdapterContract[T]](req func(savecommonglass.Input) *savecommonglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := savecommonglass.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos alterados de um vidro temperado no repositório
func PatchCommonGlasses[T contracts.FiberAdapterContract[T]](req func(patchcommonglass.Input) *patchcommonglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := patchcommonglass.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por deletar o vidro comum no repositório
func DeleteCommonGlass[T contracts.FiberAdapterContract[T]](req func(deletecommonglass.Input) *deletecommonglass.Output, ctx T) func(T) error {
	return func(c T) error {
		input := deletecommonglass.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Parts Section //

// Adaptador do fiber responsável pela obtenção de dados da peça no
// repositório de acordo com o parâmetro e pelo retorno dos mesmos para o cliente
func GetPart[T contracts.FiberAdapterContract[T]](req func(getpart.Input) *getpart.Output, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		i := getpart.Input{ID: p}
		output := req(i)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por trazer ao cliente todos os dados das peças
// comuns no repositório
func GetParts[T contracts.FiberAdapterContract[T]](req func() *getparts.Output, ctx T) func(T) error {
	return func(c T) error {
		output := req()
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por salvar a peça no repositório
func SavePart[T contracts.FiberAdapterContract[T]](req func(savepart.Input) *savepart.Output, ctx T) func(T) error {
	return func(c T) error {
		input := savepart.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos alterados de um vidro temperado no repositório
func PatchPart[T contracts.FiberAdapterContract[T]](req func(patchpart.Input) *patchpart.Output, ctx T) func(T) error {
	return func(c T) error {
		input := patchpart.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por deletar o vidro comum no repositório
func DeletePart[T contracts.FiberAdapterContract[T]](req func(deletepart.Input) *deletepart.Output, ctx T) func(T) error {
	return func(c T) error {
		input := deletepart.Input{}
		c.BodyParser(&input)
		input.UserId = c.Locals("id").(string)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// User Section //

// Adaptador do fiber responsável pela obtenção de dados do usuário no
// repositório de acordo com o parâmetro e pelo retorno dos mesmos para o cliente
func GetUser[T contracts.FiberAdapterContract[T]](req func(getuser.Input) *getuser.Output, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		i := getuser.Input{ID: p}
		output := req(i)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por trazer ao cliente todos os dados dos
// usuários no repositório
func GetUsers[T contracts.FiberAdapterContract[T]](req func() *getusers.Output, ctx T) func(T) error {
	return func(c T) error {
		output := req()
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por salvar o usuário no repositório
func SaveUser[T contracts.FiberAdapterContract[T]](req func(saveuser.Input) *saveuser.Output, ctx T) func(T) error {
	return func(c T) error {
		input := saveuser.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos alterados de um usuário no repositório
func PatchUser[T contracts.FiberAdapterContract[T]](req func(patchuser.Input) *patchuser.Output, ctx T) func(T) error {
	return func(c T) error {
		input := patchuser.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por deletar um usuário no repositório
func DeleteUser[T contracts.FiberAdapterContract[T]](req func(deleteuser.Input) *deleteuser.Output, ctx T) func(T) error {
	return func(c T) error {
		input := deleteuser.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável pela ação de login do usuário
func Login[T contracts.FiberAdapterContract[T]](req func(login.Input) *login.Output, ctx T) func(T) error {
	return func(c T) error {
		input := login.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Sale Section //

func GetSale[T contracts.FiberAdapterContract[T]](req func(getsale.Input) *getsale.Output, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		input := getsale.Input{ID: p}

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por trazer ao cliente todos os dados das
// vendas no repositório
func GetSales[T contracts.FiberAdapterContract[T]](req func() *getsales.Output, ctx T) func(T) error {
	return func(c T) error {
		output := req()
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por salvar uma venda no repositório
func SaveSale[T contracts.FiberAdapterContract[T]](req func(savesale.Input) *savesale.Output, ctx T) func(T) error {
	return func(c T) error {
		input := savesale.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por atualizar os campos alterados de uma venda no repositório
func PatchSale[T contracts.FiberAdapterContract[T]](req func(patchsale.Input) *patchsale.Output, ctx T) func(T) error {
	return func(c T) error {
		input := patchsale.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por finalizar uma venda no repositório
func CloseSale[T contracts.FiberAdapterContract[T]](req func(closesale.Input) *closesale.Output, ctx T) func(T) error {
	return func(c T) error {
		input := closesale.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}

// Adaptador do fiber responsável por deletar uma venda no repositório
func DeleteSale[T contracts.FiberAdapterContract[T]](req func(deletesale.Input) *deletesale.Output, ctx T) func(T) error {
	return func(c T) error {
		input := deletesale.Input{}
		c.BodyParser(&input)

		output := req(input)
		return c.Status(output.Status).JSON(output)
	}
}
