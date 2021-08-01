package request

type CreateProductRequest struct {
	Name		string            		`json:"name" validate:"required"`
	Code		string            		`json:"code" validate:"required"`
	Quantity	int            			`json:"quantity" validate:"required"`
}

type UpdateProductRequest struct {
	Name		string            		`json:"name" validate:"required"`
}

type SearchParamProductRequest struct {
	Name		string            		`json:"name"`
}

type SearchParamWithPagingProductRequest struct {
	Name		string      `json:"name"`
	Page 		int 		`json:"page"`
	Limit 		int 		`json:"limit"`
}