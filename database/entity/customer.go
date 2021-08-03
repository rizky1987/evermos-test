package entity

import (
	"encoding/json"
	"evermos-test/helper"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Customer struct {
	Id           				*bson.ObjectId 			`bson:"_id"`
	Code		 				string               	`bson:"code"`
	Name		 				string               	`bson:"name"`
	CreatedAtUTC        		time.Time         		`bson:"created_at_utc"`
	CreatedAtTimezone     		time.Time         		`bson:"created_at_timezone"`
	UpdatedAtUTC        		*time.Time        		`bson:"updated_at_utc,omitempty"`
	UpdatedAtTimezone     		*time.Time        		`bson:"updated_at_timezone,omitempty"`
}

// Begin Create Validation

func (entityStruct *Customer) ValidateBeforeCreate(requestedStruct request.CreateCustomerRequest) []string {

	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()
	entityStruct.CreatedAtTimezone = helper.GenerateCurrentTimeZone(helper.DefaultTimeZone)
	// Mapping All requested Struct Field to Entity Struct
	entityStruct.MappingCreateDataToEntityStruct(requestedStruct)
	entityStruct.Id = helper.CustomerIdTest

	errorResults := []string{}

	return errorResults
}

func (entityStruct *Customer) MappingCreateDataToEntityStruct(requestedStruct request.CreateCustomerRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)


}

// End Create validation

// Begin Update Validation
func (entityStruct *Customer) ValidateBeforeUpdate(requestedStruct request.UpdateCustomerRequest) error {

	currentTimeUTC := helper.GetCurrentTimeUTC()

	entityStruct.UpdatedAtUTC = &currentTimeUTC

	entityStruct.MappingUpdateDataToEntityStruct(requestedStruct)

	return nil
}

//Fungsi ini masih dalam percobaan jd jika masih ada yg blm termappingkan silahkan mapping manual
func (entityStruct *Customer) MappingUpdateDataToEntityStruct(requestedStruct request.UpdateCustomerRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	//remove all white space in right and left string field type, make sure we doesn't save unused data

}

// End Update Validation


// Please make sure all data is a correct data before we save it to DB
func (entityStruct *Customer) ValidatebeforeSaveToDB() error {
	// Add you Validation Here

	return nil
}

// End Mutual Function For Create and Update

