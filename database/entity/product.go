package entity

import (
	"encoding/json"
	"evermos-test/helper"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Product struct {
	Id           				bson.ObjectId 			`bson:"_id,omitempty"`
	Code		 				string               	`bson:"code"`
	Name		 				string               	`bson:"name"`
	Quantity		 			int               		`bson:"quantity"`
	OnHoldQuantity		 		int               		`bson:"on_hold_quantity"`
	SoldQuantity		 		int               		`bson:"sold_quantity"`
	Price		 				int               		`bson:"price"`
	MinimalStock		 		int               		`bson:"minimal_stock"`
	CreatedAtUTC        		time.Time         		`bson:"created_at_utc"`
	CreatedAtTimezone     		time.Time         		`bson:"created_at_timezone"`
	UpdatedAtUTC        		*time.Time        		`bson:"updated_at_utc,omitempty"`
	UpdatedAtTimezone     		*time.Time        		`bson:"updated_at_timezone,omitempty"`
}

// Begin Create Validation

func (entityStruct *Product) ValidateBeforeCreate(requestedStruct request.CreateProductRequest) []string {

	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()
	entityStruct.CreatedAtTimezone = helper.GenerateCurrentTimeZone(helper.DefaultTimeZone)
	// Mapping All requested Struct Field to Entity Struct
	entityStruct.MappingCreateDataToEntityStruct(requestedStruct)

	errorResults := []string{}
	entityStruct.OnHoldQuantity = 0
	entityStruct.SoldQuantity = 0

	if entityStruct.Quantity < 1 {
		errorResults = append(errorResults, "Product's quantity must greater than zero")
	}

	return errorResults
}

func (entityStruct *Product) MappingCreateDataToEntityStruct(requestedStruct request.CreateProductRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	entityStruct.MinimalStock = requestedStruct.MinimalStock
}

// End Create validation

// Begin Update Validation
func (entityStruct *Product) ValidateBeforeUpdate(requestedStruct request.UpdateProductRequest) error {

	currentTimeUTC := helper.GetCurrentTimeUTC()

	entityStruct.UpdatedAtUTC = &currentTimeUTC

	entityStruct.MappingUpdateDataToEntityStruct(requestedStruct)

	return nil
}

//Fungsi ini masih dalam percobaan jd jika masih ada yg blm termappingkan silahkan mapping manual
func (entityStruct *Product) MappingUpdateDataToEntityStruct(requestedStruct request.UpdateProductRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	//remove all white space in right and left string field type, make sure we doesn't save unused data

}

// End Update Validation


// Please make sure all data is a correct data before we save it to DB
func (entityStruct *Product) ValidatebeforeSaveToDB() error {
	// Add you Validation Here

	return nil
}

// End Mutual Function For Create and Update

