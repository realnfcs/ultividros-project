package models

import (
	"database/sql"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/realnfcs/ultividros-project/api/domain/entities"
	"gorm.io/gorm"
)

// Essa struct provém as informações base contidas na entidade de usuários
// porém com mais enfâse nas bibliotecas usadas.
type User struct {
	ID         string       `json:"id" gorm:"primaryKey"`
	Name       string       `json:"name"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	Occupation string       `json:"occupation"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" gorm:"index"`
}

// Método para criar um uuid antes de salvar no banco de dados
func (m *User) BeforeCreate(scope *gorm.DB) error {
	id := uuid.New().String()
	if id == "" {
		return errors.New("Cannot create uuid")
	}

	m.ID = strings.Replace(id, "-", "", -1)

	return nil
}

// Método responsável por transformar o model User em entidade
func (m *User) TranformToEntity() *entities.User {
	return &entities.User{
		Id:         m.ID,
		Name:       m.Name,
		Email:      m.Email,
		Password:   m.Password,
		Occupation: m.Occupation,
	}
}

// Método responsável por transformar um Slice de Models em um Slice de entidades User
func (*User) TranformToSliceOfEntity(m []User) *[]entities.User {

	user := make([]entities.User, len(m))

	var wg sync.WaitGroup

	for i, v := range m {

		wg.Add(1)

		go func() {
			user[i].Id = v.ID
			user[i].Name = v.Name
			user[i].Email = v.Email
			user[i].Password = v.Password
			user[i].Occupation = v.Occupation

			wg.Done()
		}()

		if index := i + 1; len(m)-1 > index {

			wg.Add(1)

			go func() {
				user[i+1].Id = m[i+1].ID
				user[i+1].Name = m[i+1].Name
				user[i+1].Email = m[i+1].Email
				user[i+1].Password = m[i+1].Password
				user[i+1].Occupation = m[i+1].Occupation

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &user
}

// Método responsável por transformar a entidade User em model
func (m *User) TransformToModel(e entities.User) *User {
	return &User{
		e.Id,
		e.Name,
		e.Email,
		e.Password,
		e.Occupation,
		time.Time{},
		time.Time{},
		sql.NullTime{},
	}
}

// Método que transfoma um Slice de entidades em Slice de models Users
func (*User) TransformToSliceOfModel(e []entities.User) *[]User {

	var m []User

	var wg sync.WaitGroup

	for i, v := range e {

		wg.Add(1)

		go func() {
			m[i].ID = v.Id
			m[i].Name = v.Name
			m[i].Email = v.Email
			m[i].Password = v.Password
			m[i].Occupation = v.Occupation

			wg.Done()
		}()

		if index := i + 1; len(e)-1 > index {

			wg.Add(1)

			go func() {
				m[i+1].ID = e[i+1].Id
				m[i+1].Name = e[i+1].Name
				m[i+1].Email = e[i+1].Email
				m[i+1].Password = e[i+1].Password
				m[i+1].Occupation = e[i+1].Occupation

				wg.Done()
			}()
		}

		wg.Wait()
	}

	return &m
}
