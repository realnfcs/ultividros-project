package getusers

import "github.com/realnfcs/ultividros-project/api/domain/entities"

// Usecase Output Port responsável pelos dados que serão retornados
type Output struct {
	Data   []OutputData `json:"data"`
	Status int          `json:"-"`
	Err    string       `json:"error"`
}

// OutputData é a estrutura de dados que será retornado em um array
// no Output
type OutputData struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Occupation string `json:"occupation"`
}

func (*Output) Init(i *[]entities.User, status int, err error) *Output {

	output := make([]OutputData, len(*i))

	if i == nil {
		return &Output{output, status, err.Error()}
	}

	if err != nil {
		return &Output{output, status, err.Error()}
	}

	for i, v := range *i {
		output[i].Id = v.Id
		output[i].Name = v.Name
		output[i].Email = v.Email
		output[i].Password = v.Password
		output[i].Occupation = v.Occupation
	}

	return &Output{output, status, ""}
}
