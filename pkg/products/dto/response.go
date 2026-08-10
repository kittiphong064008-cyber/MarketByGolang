package dto

type Products struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Descript string `json:"descript"`
}

type ProductResponse struct {
	TotalItem int `json:"Total"`
	MaxPrice  int `json:"Max"`
	MinPrice  int `json:"Min"`
	Products  []Products
}

type ResponseProducts struct {
	Message string          `json:"message"`
	Value   ProductResponse `json:"products"`
}

type ResponseOne struct {
	Message string   `json:"message"`
	Value   Products `json:"products"`
}

type ResponseAll struct {
	Message string     `json:"message"`
	Value   []Products `json:"products"`
}

type ResponseRows struct {
	Message    string `json:"message"`
	RowsAffect int    `json:"RowsAffect"`
}

type ResponseError struct {
	Message string `json:"message"`
}
