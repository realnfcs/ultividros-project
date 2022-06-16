package entities

// Essa struct provém as informações contidas na entidade dos vidros comuns
type CommonGlass struct {
	Id              string  `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float32 `json:"price"`
	Type            string  `json:"type"`
	Color           string  `json:"color"`
	Milimeter       float32 `json:"milimeter"`
	HeightAvailable float32 `json:"height_available"`
	WidthAvailable  float32 `json:"width_available"`
}
