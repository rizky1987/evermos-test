package entity

import (
	"encoding/json"
	"evermos-test/helper"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"time"
)

type HomeBlock struct {
	Id           				*bson.ObjectId 			`bson:"_id,omitempty"`
	Block		 				string               	`bson:"block"`
	Sequence		 			int               		`bson:"sequence"`
	CreatedAt        			time.Time         		`bson:"created_at"`
	CreatedAtUTC     			time.Time         		`bson:"created_at_utc"`
	UpdatedAt        			*time.Time        		`bson:"updated_at,omitempty"`
	UpdatedAtUTC     			*time.Time        		`bson:"updated_at_utc,omitempty"`
}

// Begin Create Validation

func (entityStruct *HomeBlock) ValidateBeforeCreate(requestedStruct request.CreateHomeBlockRequest) error {

	entityStruct.CreatedAt = helper.GetCurrentTimeAsiaJakarta()
	entityStruct.CreatedAtUTC = helper.GetCurrentTimeUTC()

	// Mapping All requested Struct Field to Entity Struct
	entityStruct.MappingCreateDataToEntityStruct(requestedStruct)

	return nil
}

//Fungsi ini masih dalam percobaan jd jika masih ada yg blm termappingkan silahkan mapping manual
func (entityStruct *HomeBlock) MappingCreateDataToEntityStruct(requestedStruct request.CreateHomeBlockRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	//remove all white space in right and left string field type, make sure we doesn't save unused data
	entityStruct.TrimStringFieldType()

}

// End Create validation

// Begin Update Validation
func (entityStruct *HomeBlock) ValidateBeforeUpdate(requestedStruct request.UpdateHomeBlockRequest) error {

	currentTimeAsiaJakarta := helper.GetCurrentTimeAsiaJakarta()
	currentTimeUTC := helper.GetCurrentTimeUTC()

	entityStruct.UpdatedAt = &currentTimeAsiaJakarta
	entityStruct.UpdatedAtUTC = &currentTimeUTC

	entityStruct.MappingUpdateDataToEntityStruct(requestedStruct)

	return nil
}

//Fungsi ini masih dalam percobaan jd jika masih ada yg blm termappingkan silahkan mapping manual
func (entityStruct *HomeBlock) MappingUpdateDataToEntityStruct(requestedStruct request.UpdateHomeBlockRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	//remove all white space in right and left string field type, make sure we doesn't save unused data
	entityStruct.TrimStringFieldType()

}

// End Update Validation

// Begin Mutual Function For Create and Update
func (entityStruct *HomeBlock) TrimStringFieldType() {

	msValuePtr := reflect.ValueOf(entityStruct)
	msValue := msValuePtr.Elem()

	for i := 0; i < msValue.NumField(); i++ {
		field := msValue.Field(i)

		// Ignore fields that don't have the same type as a string
		if field.Type() != reflect.TypeOf("") {
			continue
		}

		str := field.Interface().(string)
		str = helper.TrimWhiteSpace(str)
		field.SetString(str)
	}
}

// Please make sure all data is a correct data before we save it to DB
func (entityStruct *HomeBlock) ValidatebeforeSaveToDB() error {

	entityStruct.TrimStringFieldType()
	// Add you Validation Here

	return nil
}

// End Mutual Function For Create and Update

