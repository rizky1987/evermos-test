package response

import (
	"encoding/json"
	"evermos-test/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type CartResponse struct {
	Id          		bson.ObjectId   		`json:"id"`
	ProductId			string            		`json:"product_id"`
	CustomerId			string            		`json:"customer_id"`
	Quantity		 	int               		`json:"quantity"`
}

type CartSuccessResponse struct {
	CommonBaseResponse
	Data 			CartResponse 		`json:"result"`
}

type CartSuccessWithPagingResponse struct {
	CommonBaseResponse
	Data 			*CartSearchResponse `json:"result"`
}

type CartSearchResponse struct {
	Data 			[]*CartResponse 			`json:"data"`
	CurrentPage 	int 						`json:"current_page"`
	Limit 			int 						`json:"limit"`
	TotalRecords 	int 						`json:"total_records"`
	TotalPages 		int 						`json:"total_page"`
}


func (output *CartSearchResponse) GeneratePagingResponse(data []*CartResponse, page, limit, totalRecords int ) {

	output.Data = data
	output.CurrentPage = page
	output.Limit = limit
	output.TotalRecords =  totalRecords
	output.TotalPages = 1
	if limit > 0 {
		output.TotalPages = (totalRecords / limit)
	}
}

type CartFailedResponse struct {
	CommonBaseResponse
}

func (output *CartResponse) ParsingEntityToResponse(inputedEntity *entity.Cart) {

	// fungsi ini masih dalam percaobaan jika ada yang tidak terparsing, silahkan parsing manual. :D

	jsonString, _ := json.Marshal(inputedEntity)
	json.Unmarshal(jsonString, &output)

}
