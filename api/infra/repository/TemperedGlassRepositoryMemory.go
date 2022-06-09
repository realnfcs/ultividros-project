package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/realnfcs/ultividros-project/api/domain/entities"
)

// Estrutura de dados (Struct) para o armazenamento de dados em memória
type TemperedGlassRepositoryMemory struct {
	TemperedGlasses []entities.TemperedGlass
}

// Função para iniciar o repositório em memória
func (*TemperedGlassRepositoryMemory) Init() *TemperedGlassRepositoryMemory {
	return &TemperedGlassRepositoryMemory{
		[]entities.TemperedGlass{
			{
				Id:          "184751d5-fa3f-4837-b3de-4831164aac1b",
				Name:        "Janela",
				Description: "Uma janela temperada de 4 folhas",
				Price:       200,
				Quantity:    5,
				Type:        "Tempered",
				Color:       "smoked",
				GlassSheets: 4,
				Milimeter:   8,
				Height:      1.20,
				Width:       1.0,
			},
			{
				Id:          "b220daae-f3e4-4aa4-9cd5-be00dcb1325f",
				Name:        "Janela",
				Description: "Uma janela temperada de 2 folhas",
				Price:       200,
				Quantity:    5,
				Type:        "Tempered",
				Color:       "Transparent",
				GlassSheets: 2,
				Milimeter:   10,
				Height:      1.5,
				Width:       1.2,
			},
			{
				Id:          "67c3ff2f-3961-47c1-acbb-4101e5687008",
				Name:        "Fixo",
				Description: "Um fixo temperado",
				Price:       130,
				Quantity:    3,
				Type:        "Tempered",
				Color:       "Transparent",
				GlassSheets: 1,
				Milimeter:   10,
				Height:      2,
				Width:       1.2,
			},
		},
	}
}

// Método que retorna os dados do vidro temperado no repositório em memório
// de acordo com o ID passado no parâmetro
func (t *TemperedGlassRepositoryMemory) GetTemperedGlass(id string) *entities.TemperedGlass {
	for i, v := range t.TemperedGlasses {
		if v.Id == id {
			return &t.TemperedGlasses[i]
		}
	}

	return nil
}

// Método que retorna todos os dados armazenado no repositório em memória
func (t *TemperedGlassRepositoryMemory) GetTemperedGlasses() *[]entities.TemperedGlass {
	return &t.TemperedGlasses
}

// Método que salva o vidro temperado no repositório em memória
func (t *TemperedGlassRepositoryMemory) SaveTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	e.Id = uuid.New().String()
	t.TemperedGlasses = append(t.TemperedGlasses, e)
	return e.Id, 201, nil
}

// Método que atualiza o vidro temperado no repositório em memória
func (t *TemperedGlassRepositoryMemory) UpdateTemperedGlass(e entities.TemperedGlass) (string, int, error) {
	temperedGlass := t.GetTemperedGlass(e.Id)
	if temperedGlass == nil {
		return "", 401, errors.New("a nil pointer is received: this glass in this id not exist")
	}

	ent := entities.TemperedGlass{}

	for _, v := range t.TemperedGlasses {
		if v.Id == e.Id {

			ent.Id = e.Id
			ent.Name = e.Name

			ent.Description = e.Description
			ent.Price = e.Price
			ent.Quantity = e.Quantity
			ent.Type = e.Type
			ent.Color = e.Color
			ent.GlassSheets = e.GlassSheets
			ent.Milimeter = e.Milimeter
			ent.Height = e.Height
			ent.Width = e.Width

			t.TemperedGlasses = removeIt(v, t.TemperedGlasses)
			t.TemperedGlasses = append(t.TemperedGlasses, ent)

			return ent.Id, 401, nil
		}
	}

	return "", 401, errors.New("cannot save the tempered glass")
}

func removeIt(e entities.TemperedGlass, t []entities.TemperedGlass) []entities.TemperedGlass {
	for i, v := range t {
		if v == e {
			return append(t[0:i], t[i+1:]...)
		}
	}
	return t
}
