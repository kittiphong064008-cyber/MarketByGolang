package dto

type Products struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Descript string `json:"descript"`
}

type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
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

type ResponsePaginated struct {
	Message    string     `json:"message"`
	Value      []Products `json:"products"`
	Pagination Pagination `json:"pagination"`
}

type ResponseRows struct {
	Message    string `json:"message"`
	RowsAffect int    `json:"RowsAffect"`
}

type ResponseError struct {
	Message string `json:"message"`
}

func BuildPagination(page, limit, total int) *Pagination {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	if total <= 0 {
		total = 0
	}

	totalPages := total / limit
	if total%limit != 0 {
		totalPages++
	}
	if totalPages == 0 {
		totalPages = 1
	}

	return &Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPages,
	}
}
