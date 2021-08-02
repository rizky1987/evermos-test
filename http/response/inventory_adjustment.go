package response

import (
	"encoding/json"
	"evermos-test/database/entity"
	"gopkg.in/mgo.v2/bson"
)

type InventoryAdjustmentResponse struct {
	Id          		bson.ObjectId   		`json:"id"`
	Code				string            		`json:"code"`
	Name				string            		`json:"name"`
	Quantity		 	int               		`json:"quantity"`
	OnHoldQuantity		int               		`json:"on_hold_quantity"`
	SoldQuantity		int               		`json:"sold_quantity"`
	Message				string					`json:"message"`
}

type InventoryAdjustmentSuccessResponse struct {
	CommonBaseResponse
	Data 			InventoryAdjustmentResponse 		`json:"result"`
}

type InventoryAdjustmentSuccessWithPagingResponse struct {
	CommonBaseResponse
	Data 			*InventoryAdjustmentSearchResponse `json:"result"`
}

type InventoryAdjustmentSearchResponse struct {
	Data 			[]*InventoryAdjustmentResponse 			`json:"data"`
	CurrentPage 	int 						`json:"current_page"`
	Limit 			int 						`json:"limit"`
	TotalRecords 	int 						`json:"total_records"`
	TotalPages 		int 						`json:"total_page"`
}


func (output *InventoryAdjustmentSearchResponse) GeneratePagingResponse(data []*InventoryAdjustmentResponse, page, limit, totalRecords int ) {

	output.Data = data
	output.CurrentPage = page
	output.Limit = limit
	output.TotalRecords =  totalRecords
	output.TotalPages = 1
	if limit > 0 {
		output.TotalPages = (totalRecords / limit)
	}
}

type InventoryAdjustmentFailedResponse struct {
	CommonBaseResponse
}

func (output *InventoryAdjustmentResponse) ParsingEntityToResponse(inputedEntity *entity.InventoryAdjustment) {

	jsonString, _ := json.Marshal(inputedEntity)
	json.Unmarshal(jsonString, &output)
}
