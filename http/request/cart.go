package request

type CreateCartRequest struct {
	ProductId		string            		`json:"product_id" validate:"required"`
	CustomerId		string            		`json:"customer_id" validate:"required"`
	Quantity		int        				`json:"quantity" validate:"required"`
}

type CheckoutCartRequest struct {
	Ids			    []string            	`json:"cart_ids" validate:"required"`
}

type UpdateCartRequest struct {
	ProductId		string            		`json:"product_id" validate:"required"`
}

type SearchParamCartRequest struct {
	Name		string            			`json:"name"`
}

type SearchParamWithPagingCartRequest struct {
	Name		string      `json:"name"`
	Page 		int 		`json:"page"`
	Limit 		int 		`json:"limit"`
}