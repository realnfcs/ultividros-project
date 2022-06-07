package adapters

import (
	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"github.com/realnfcs/ultividros-project/api/interface/contracts"
)

// Adaptador do fiber responsável pela obtenção de dados do vidro temperado no
// repositório de acordo com o parâmetro e pelo retorno dos mesmos para o cliente
func GetTemperedGlass[T contracts.Adapters](req func(string) *entities.TemperedGlass, ctx T) func(T) error {
	return func(c T) error {
		p := c.Params("id")
		obj := req(p)
		return c.JSON(obj)
	}
}

// Adaptador do fiber responsável por trazer ao cliente todos os dados dos vidros
// temperados no repositório
func GetTemperedGlasses[T contracts.Adapters](req func() *[]entities.TemperedGlass, ctx T) func(T) error {
	return func(c T) error {
		obj := req()
		return c.JSON(obj)
	}
}

// Adaptador do fiber responsável por salvar o vidro temperado no repositório
func SaveTemperedGlasses[T contracts.Adapters](req func(entities.TemperedGlass) error, ctx T) func(T) error {
	return func(c T) error {
		ent := entities.TemperedGlass{}
		c.BodyParser(&ent)

		err := req(ent)
		if err != nil {
			return err
		}

		return c.SendStatus(200)
	}
}

func UpdateTemperedGlasses[T contracts.Adapters](req func(entities.TemperedGlass) error, ctx T) func(T) error {
	return func(c T) error {
		ent := entities.TemperedGlass{}
		c.BodyParser(&ent)

		err := req(ent)
		if err != nil {
			return err
		}

		return c.SendStatus(200)
	}
}
