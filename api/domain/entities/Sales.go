package entities

// Essa struct provém as informações contidas na entidade das vendas
// Cada produto requisito pelo cliente será armazendo nesta struct
// separado por categorias/entities que participam na venda, são eles:
// 		- Vidros comuns (CommonGlssReq)
//		- Vidros temperados (TempGlssReq)
//	 	- Peças (PartsReq)
// Cada produto seja qualquer um citado a cima corresponde a um produto
// do repositório porém destinado para a(s) venda(s). Ou seja:
//		- CommonGlssReq próvem as informações dos vidros comuns
//		  requeridos na venda pelo cliente e assim funciona também para
//  	  os outros.
type Sale struct {
	Id            string          `json:"id"`
	ClientId      string          `json:"client_id"`
	CommonGlssReq []CommonGlssReq `json:"common_glss_req"`
	PartsReq      []PartsReq      `json:"parts_req"`
	TempGlssReq   []TempGlssReq   `json:"temp_glss"`
	IsActive      bool            `json:"is_active"`
}

// Essa struct provém as informações dos produtos requisitados na venda,
// representando os vidros temperados, comuns e peças que o cliente poderá comprar
type ProductsRequest struct {
	Id           string  `json:"id"`
	ProductId    string  `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float32 `json:"product_price"`
	ProdtQtyReq  uint32  `json:"prodt_qty_req"`
	WasCancelled bool    `json:"was_cancelled"`
	WasConfirmed bool    `json:"was_confimed"`
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
