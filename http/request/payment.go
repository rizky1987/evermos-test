package request

type CallbackPaymentRequest struct {
	PaymentCode		string            		`json:"payment_code" validate:"required"`
	Status			string            		`json:"status" validate:"required"`
}

type SearchParamPaymentRequest struct {
	PaymentCode		string            		`json:"payment_code"`
}

type SearchParamWithPagingPaymentRequest struct {
	PaymentCode		string      			`json:"payment_code"`
	Page 			int 					`json:"page"`
	Limit 			int 					`json:"limit"`
}