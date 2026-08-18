package dto

type CatagoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseOne struct {
	Message string           `json:"message"`
	Value   CatagoryResponse `json:"products"`
}

type ResponseAll struct {
	Message string             `json:"message"`
	Value   []CatagoryResponse `json:"products"`
}

type ResponseRows struct {
	Message    string `json:"message"`
	RowsAffect int    `json:"RowsAffect"`
}

type ResponseError struct {
	Message string `json:"message"`
}
