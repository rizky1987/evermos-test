package entity

import (
	"evermos-test/helper"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type InventoryAdjustment struct {
	Id                		*bson.ObjectId  `bson:"_id,omitempty"`
	ProductId         		*bson.ObjectId  `bson:"product_id"`
	Process           		string         `bson:"process"`
	Quantity          		int            `bson:"quantity"`
	Note    				string         `bson:"note"`
	CreatedAtUTC      		time.Time      `bson:"created_at_utc"`
	CreatedAtTimezone 		time.Time      `bson:"created_at_timezone"`
	UpdatedAtUTC      		*time.Time     `bson:"updated_at_utc,omitempty"`
	UpdatedAtTimezone 		*time.Time     `bson:"updated_at_timezone,omitempty"`
}

// Begin Create Validation

func (entityStruct *InventoryAdjustment) ValidateBeforeCreate() []string {

	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()
	entityStruct.CreatedAtTimezone = helper.GenerateCurrentTimeZone(helper.DefaultTimeZone)

	errorResults := []string{}

	if entityStruct.Quantity < 1 {
		errorResults = append(errorResults, "Inventory Adjustment's quantity must greater than zero")
	}
	return errorResults
}

// End Create validation