package entity

import (
	"encoding/json"
	"evermos-test/helper"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type InventoryAdjustment struct {
	Id                		*bson.ObjectId  `bson:"_id,omitempty"`
	ProductId         		*bson.ObjectId  `bson:"product_id"`
	Process           		string        	`bson:"process"`
	Quantity          		int           	`bson:"quantity"`
	Note    				string        	`bson:"note"`
	CreatedAtUTC      		time.Time     	`bson:"created_at_utc"`
	CreatedAtTimezone 		time.Time     	`bson:"created_at_timezone"`
}

// Begin Create Validation

func (entityStruct *InventoryAdjustment) ValidateBeforeCreate() []string {

	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()
	entityStruct.CreatedAtTimezone = helper.GenerateCurrentTimeZone(helper.DefaultTimeZone)

	errorResults := []string{}

	return errorResults
}

// End Create validation

func (entityStruct *InventoryAdjustment) MappingCreateDataToEntityStruct(requestedStruct request.CreateInventoryAdjustmentRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	productId, _ := helper.ChangeStringOfObjectIdToBsonObjectId(requestedStruct.ProductId)
	entityStruct.ProductId = productId
}