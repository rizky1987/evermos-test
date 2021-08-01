package response

type CommonBaseResponse struct {
	Message  []string `json:"error_messages"`
	Code     int      `json:"code"`
	CodeType string   `json:"code_type"`
}

type CommonPagingResponse struct {
	CurrentPage 	int 			`json:"current_page"`
	Limit 			*int 			`json:"limit"`
	TotalRecords 	*int 			`json:"total_records"`
	TotalPages 		*int 			`json:"total_page"`
}
