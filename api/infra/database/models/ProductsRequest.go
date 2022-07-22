package models

import (
	"database/sql"
	"time"
)

// Essa struct provém as informações dos produtos requisitados na venda,
// representando os vidros temperados, comuns e peças que o cliente poderá comprar
type ProductsRequest struct {
	ID           string       `json:"id" gorm:"primaryKey"`
	ProductId    string       `json:"product_id"`
	ProductName  string       `json:"product_name"`
	ProductPrice float32      `json:"product_price"`
	ProdtQtyReq  uint32       `json:"prodt_qty_req"`
	WasCancelled bool         `json:"was_cancelled"`
	WasConfirmed bool         `json:"was_confimed"`
	SaleID       string       `json:"sale_id" gorm:"size:191"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" gorm:"index"`
}
