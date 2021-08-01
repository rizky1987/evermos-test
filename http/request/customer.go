package request

type CreateCustomerRequest struct {
	Name		string            		`json:"name" validate:"required"`
	Code		string            		`json:"code" validate:"required"`
}

type UpdateCustomerRequest struct {
	Name		string            		`json:"name" validate:"required"`
}

type SearchParamCustomerRequest struct {
	Name		string            		`json:"name"`
}

type SearchParamWithPagingCustomerRequest struct {
	Name		string      `json:"name"`
	Page 		int 		`json:"page"`
	Limit 		int 		`json:"limit"`
}