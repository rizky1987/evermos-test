package response

import (
	"encoding/json"
	"evermos-test/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type PaymentResponse struct {
	Id          		bson.ObjectId   		`json:"id"`
	Code				string            		`json:"code"`
	Name				string            		`json:"name"`
	Quantity		 	int               		`json:"quantity"`
	OnHoldQuantity		int               		`json:"on_hold_quantity"`
	SoldQuantity		int               		`json:"sold_quantity"`
	Message				string					`json:"message"`
}

type PaymentSuccessResponse struct {
	CommonBaseResponse
	Data 			PaymentResponse 		`json:"result"`
}

type PaymentSuccessWithPagingResponse struct {
	CommonBaseResponse
	Data 			*PaymentSearchResponse `json:"result"`
}

type PaymentSearchResponse struct {
	Data 			[]*PaymentResponse 			`json:"data"`
	CurrentPage 	int 						`json:"current_page"`
	Limit 			int 						`json:"limit"`
	TotalRecords 	int 						`json:"total_records"`
	TotalPages 		int 						`json:"total_page"`
}


func (output *PaymentSearchResponse) GeneratePagingResponse(data []*PaymentResponse, page, limit, totalRecords int ) {

	output.Data = data
	output.CurrentPage = page
	output.Limit = limit
	output.TotalRecords =  totalRecords
	output.TotalPages = 1
	if limit > 0 {
		output.TotalPages = (totalRecords / limit)
	}
}

type PaymentFailedResponse struct {
	CommonBaseResponse
}

func (output *PaymentResponse) ParsingEntityToResponse(inputedEntity *entity.Payment) {

	jsonString, _ := json.Marshal(inputedEntity)
	json.Unmarshal(jsonString, &output)
}
