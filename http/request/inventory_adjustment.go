package request

type CreateInventoryAdjustmentRequest struct {
	ProductId         		string  		`json:"product_id" validate:"required"`
	Process           		string        	`json:"process" validate:"required"`
	Quantity          		int           	`json:"quantity" validate:"required"`
	Note    				string        	`json:"note" validate:"required"`
}

type SearchParamInventoryAdjustmentRequest struct {
	ProductId		string            		`json:"product_id"`
}

type SearchParamWithPagingInventoryAdjustmentRequest struct {
	ProductId		string            		`json:"product_id"`
	Page 			int 					`json:"page"`
	Limit 			int 					`json:"limit"`
}