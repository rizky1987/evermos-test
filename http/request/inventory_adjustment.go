package request

type SearchParamInventoryAdjustmentRequest struct {
	ProductName		string            		`json:"product_name"`
	ProductCode		string            		`json:"product_code"`
}

type SearchParamWithPagingInventoryAdjustmentRequest struct {
	ProductName		string            		`json:"product_name"`
	ProductCode		string            		`json:"product_code"`
	Page 			int 					`json:"page"`
	Limit 			int 					`json:"limit"`
}