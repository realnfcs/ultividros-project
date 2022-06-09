// Pacote responsável pela inicialização do banco de dados MySQL
package databasemysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Struct responsável pela o armazenamento dos campos do dsn para a conexão
type DatabaseMysql struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DbName string
}

// Método para iniciar a conexão com o Banco de Dados de acordo com as variáveis
// de ambiente contidas no arquivo .env
func (d *DatabaseMysql) Init() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	d.User = os.Getenv("DB_USER")
	d.Pass = os.Getenv("DB_PASS")
	d.Host = os.Getenv("DB_HOST")
	d.Port = os.Getenv("DB_PORT")
	d.DbName = os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", d.User, d.Pass, d.Host, d.Port, d.DbName)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}
