package contracts

// Interface para que bibliotecas de servidor HTTP externas (Frameworks and Drivers Layer)
// sejam usadas pelos adaptdores para que ocorra a inversão de dependência
type Adapters interface {
	JSON(any) error
	BodyParser(out any) error
	Params(key string, defaultValue ...string) string
	SendStatus(status int) error
}
