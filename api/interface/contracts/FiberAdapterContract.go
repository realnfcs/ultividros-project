package contracts

// Interface para que biblioteca de servidor HTTP externa (gofiber)
// seja usada pelo adaptador para que ocorra a inversão de dependência
type FiberAdapterContract[T any] interface {
	Adapters
	Status(int) T
	Locals(key string, value ...interface{}) interface{}
}
