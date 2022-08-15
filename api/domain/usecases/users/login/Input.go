package login

// Usecase Input Port responsável pelos dados que entrarão
type Input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
