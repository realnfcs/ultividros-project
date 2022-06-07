// Pacote responsável pela o usecase GetTemperedGlasses que executa
// a ação de pegar todos os vidro temperados e retornar os dados ao
// cliente
package gettemperedglasses

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Data []OutputData
}

// OutputData é a estrutura de dados que será retornado em um array
// no Output
type OutputData struct {
	Id          string  `json:"id"`
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
