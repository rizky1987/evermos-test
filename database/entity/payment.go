package entity

import (
	"evermos-test/helper"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Payment struct {
	Id           				bson.ObjectId 			`bson:"_id,omitempty"`
	Code		 				string               	`bson:"code,omitempty"`
	Total		 				int               		`bson:"total,omitempty"`
	Status		 				string               	`bson:"status,omitempty"`
	CreatedAtUTC        		time.Time         		`bson:"created_at_utc"`
	CreatedAtTimezone     		time.Time         		`bson:"created_at_timezone"`
	UpdatedAtUTC        		*time.Time        		`bson:"updated_at_utc,omitempty"`
	UpdatedAtTimezone     		*time.Time        		`bson:"updated_at_timezone,omitempty"`
}

// Begin Create Validation

func (entityStruct *Payment) ValidateBeforeCreate() []string {

	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()
	entityStruct.CreatedAtTimezone = helper.GenerateCurrentTimeZone(helper.DefaultTimeZone)

	errorResults := []string{}
	return errorResults
}

// End Create validation

// End Update Validation

// Please make sure all data is a correct data before we save it to DB
func (entityStruct *Payment) ValidatebeforeSaveToDB() error {
	// Add you Validation Here

	return nil
}

// End Mutual Function For Create and Update

