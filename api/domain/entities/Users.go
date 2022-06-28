package entities

// Essa struct provém as informações contidas na entidade dos usuários
type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Occupation string `json:"occupation"`
}
