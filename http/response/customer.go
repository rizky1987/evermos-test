package response

import (
	"encoding/json"
	"evermos-test/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type CustomerResponse struct {
	Id          bson.ObjectId   		`json:"id"`
	Name		string            		`json:"name" validate:"required"`
	Code		string            		`json:"code" validate:"required"`
}

type CustomerSuccessResponse struct {
	CommonBaseResponse
	Data 			CustomerResponse 		`json:"result"`
}

type CustomerSuccessWithPagingResponse struct {
	CommonBaseResponse
	Data 			*CustomerSearchResponse `json:"result"`
}

type CustomerSearchResponse struct {
	Data 			[]*CustomerResponse 			`json:"data"`
	CurrentPage 	int 						`json:"current_page"`
	Limit 			int 						`json:"limit"`
	TotalRecords 	int 						`json:"total_records"`
	TotalPages 		int 						`json:"total_page"`
}


func (output *CustomerSearchResponse) GeneratePagingResponse(data []*CustomerResponse, page, limit, totalRecords int ) {

	output.Data = data
	output.CurrentPage = page
	output.Limit = limit
	output.TotalRecords =  totalRecords
	output.TotalPages = 1
	if limit > 0 {
		output.TotalPages = (totalRecords / limit)
	}
}

type CustomerFailedResponse struct {
	CommonBaseResponse
}

func (output *CustomerResponse) ParsingEntityToResponse(inputedEntity *entity.Customer) {

	// fungsi ini masih dalam percaobaan jika ada yang tidak terparsing, silahkan parsing manual. :D

	jsonString, _ := json.Marshal(inputedEntity)
	json.Unmarshal(jsonString, &output)

}
