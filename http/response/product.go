package response

import (
	"encoding/json"
	"evermos-test/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type ProductResponse struct {
	Id          		bson.ObjectId   		`json:"id"`
	Code				string            		`json:"code"`
	Name				string            		`json:"name"`
	Quantity		 	int               		`json:"quantity"`
	OnHoldQuantity		int               		`json:"on_hold_quantity"`
	SoldQuantity		int               		`json:"sold_quantity"`
}

type ProductSuccessResponse struct {
	CommonBaseResponse
	Data 			ProductResponse 		`json:"result"`
}

type ProductSuccessWithPagingResponse struct {
	CommonBaseResponse
	Data 			*ProductSearchResponse `json:"result"`
}

type ProductSearchResponse struct {
	Data 			[]*ProductResponse 			`json:"data"`
	CurrentPage 	int 						`json:"current_page"`
	Limit 			int 						`json:"limit"`
	TotalRecords 	int 						`json:"total_records"`
	TotalPages 		int 						`json:"total_page"`
}


func (output *ProductSearchResponse) GeneratePagingResponse(data []*ProductResponse, page, limit, totalRecords int ) {

	output.Data = data
	output.CurrentPage = page
	output.Limit = limit
	output.TotalRecords =  totalRecords
	output.TotalPages = 1
	if limit > 0 {
		output.TotalPages = (totalRecords / limit)
	}
}

type ProductFailedResponse struct {
	CommonBaseResponse
}

func (output *ProductResponse) ParsingEntityToResponse(inputedEntity *entity.Product) {

	// fungsi ini masih dalam percaobaan jika ada yang tidak terparsing, silahkan parsing manual. :D

	jsonString, _ := json.Marshal(inputedEntity)
	json.Unmarshal(jsonString, &output)

}
