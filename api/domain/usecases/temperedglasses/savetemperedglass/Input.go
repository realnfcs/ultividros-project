// Pacote responsável pela o usecase SaveTemperedGlass que executa
// a ação de salvamento de um vidro temperado e retorna os dado de
// resposta ao cliente
package savetemperedglass

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32  `json:"quantity"`
	Type        string  `json:"type"`
	Color       string  `json:"color"`
	GlassSheets uint8   `json:"glass_sheets"`
	Milimeter   float32 `json:"milimeter"`
	Height      float32 `json:"height"`
	Width       float32 `json:"width"`
}
