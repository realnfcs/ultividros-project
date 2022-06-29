package controllers

import (
	"github.com/realnfcs/ultividros-project/api/domain/repository"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/deleteuser"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/getuser"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/getusers"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/patchuser"
	"github.com/realnfcs/ultividros-project/api/domain/usecases/users/saveuser"
)

// Stuct para iniciar a controller dos usuários necessitando
// de um repository para funcionar
type UserController struct {
	Repo repository.UserRepository
}

// Método da controller que comunica com o usecase para a obtenção de dados de uma
// única peça trazendo a resposta ao cliente
func (u *UserController) GetUser(i getuser.Input) *getuser.Output {
	getUser := getuser.GetUser{UserRepository: u.Repo}
	output := getUser.Execute(i)
	return output
}

// Método da controller que comunica com o usecase para a obtenção de dados e pela
// resposta ao cliente
func (u *UserController) GetUsers() *getusers.Output {
	getUsers := getusers.GetUsers{UserRepository: u.Repo}
	output := getUsers.Execute()
	return output
}

// Método da controller que comunica com o usecase para salvar um objeto usuário
// de acordo com os dados passados no parâmetro
func (u *UserController) SaveUser(i saveuser.Input) *saveuser.Output {
	saveUser := saveuser.SaveUser{UserRepository: u.Repo}
	output := saveUser.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para atualizar os campos alterados do objeto de
// acordo com os dados passados no parâmetro
func (u *UserController) PatchUser(i patchuser.Input) *patchuser.Output {
	patchUser := patchuser.PatchUser{UserRepository: u.Repo}
	output := patchUser.Execute(i)
	return output
}

// Método da controller que comunica com o usacase para deletar um objeto
func (u *UserController) DeleteUser(i deleteuser.Input) *deleteuser.Output {
	deleteUser := deleteuser.DeleteUser{UserRepository: u.Repo}
	output := deleteUser.Execute(i)
	return output
}
