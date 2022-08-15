// Pacote responsável pela o usecase Login que executa
// a ação de verificar o email e senha do usuário e volta
// um token para o cliente proporcionando a ação do login
package login

import "github.com/realnfcs/ultividros-project/api/domain/repository"

// Usecase responsável pela a ação do login do usuário recebendo
// um email e password e voltando um token para o cliente
type Login struct {
	UserRepository repository.UserRepository
}

func (l *Login) Execute(i Input) *Output {
	token, status, err := l.UserRepository.Login(i.Email, i.Password)
	return new(Output).Init(token, status, err)
}
