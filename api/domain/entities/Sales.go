package entities

// Essa struct provém as informações contidas na entidade das vendas
// Obs.: o campo ProductsId guarda os IDs das requisições de produtos
// feitos pelos cliente, provendo melhores querys.
// Ex.: ProductsId["ID1", "ID2", "ID3", ...]
// 		Cada ID aponta para uma tabela com uma requisição de produto criada
//		pelo clinte quando esse os requisita, desta forma:
//		  "ID1" --> ProductsRequest { "ID1", "ProductID1", "Produto A exemplo", ... }
//		  "ID2" --> ProductsRequest { "ID2", "ProductID2", "Produto B exemplo", ... }
// 		E cada ProductID aponta para um vidro comum ou temperado ou para uma peça
type Sale struct {
	Id          string   `json:"id"`
	ClientId    string   `json:"client_id"`
	ProductsIds []string `json:"products_ids"`
}

// Essa struct provém as informações dos produtos requisitados na venda,
// representando os vidros temperados, comuns e peças que o cliente poderá comprar
type ProductsRequest struct {
	Id           string  `json:"id"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	ProdtQtyReq  uint32  `json:"prodt_qty_req"`
	SaleId       string  `json:"sale_id"`
}

type TempGlssReq struct {
	ProductsRequest
}

type PartsReq struct {
	ProductsRequest
}

// Essa struct provém as informações mais específicas dos vidros comuns
// requisitados na venda, por ter uma estrutura de dados diferentes das
// outras, não sendo possível generaliza-lo com uma struct como acontece
// com os outros produtos
type CommonGlssReq struct {
	ProductsRequest
	RequestWidth  float32 `json:"request_width"`
	RequestHeight float32 `json:"request_height"`
}
