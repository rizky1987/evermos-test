package entity

import (
	"encoding/json"
	"evermos-test/helper"
	"evermos-test/http/request"
	"gopkg.in/mgo.v2/bson"
)

type Cart struct {
	Id           				*bson.ObjectId 			`bson:"_id,omitempty"`
	ProductId		 			*bson.ObjectId          `bson:"product_id"`
	CustomerId		 			*bson.ObjectId          `bson:"customer_id"`
	Quantity		 			int               		`bson:"quantity"`
	Status		 				string               	`bson:"status"`
}

// Begin Create Validation

func (entityStruct *Cart) ValidateBeforeCreate(requestedStruct request.CreateCartRequest) []string {

	// Mapping All requested Struct Field to Entity Struct
	entityStruct.MappingCreateDataToEntityStruct(requestedStruct)

	errorResults := []string{}

	productId, errProductId := helper.ChangeStringOfObjectIdToBsonObjectId(requestedStruct.ProductId)
	if errProductId != "" {
		errorResults = append(errorResults, helper.ErrorIsNOTObjectIdHex(requestedStruct.ProductId))
	}

	customerId, errCustomerId := helper.ChangeStringOfObjectIdToBsonObjectId(requestedStruct.CustomerId)
	if errCustomerId != "" {
		errorResults = append(errorResults, helper.ErrorIsNOTObjectIdHex(requestedStruct.CustomerId))
	}

	entityStruct.Status = helper.CartStatusNew
	entityStruct.ProductId = productId
	entityStruct.CustomerId = customerId
	return errorResults
}

func (entityStruct *Cart) MappingCreateDataToEntityStruct(requestedStruct request.CreateCartRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)
}

// End Create validation

// Begin Update Validation
func (entityStruct *Cart) ValidateBeforeUpdate(requestedStruct request.UpdateCartRequest) error {

	entityStruct.MappingUpdateDataToEntityStruct(requestedStruct)

	return nil
}

//Fungsi ini masih dalam percobaan jd jika masih ada yg blm termappingkan silahkan mapping manual
func (entityStruct *Cart) MappingUpdateDataToEntityStruct(requestedStruct request.UpdateCartRequest) {

	jsonString, _ := json.Marshal(requestedStruct)
	json.Unmarshal(jsonString, &entityStruct)

	//remove all white space in right and left string field type, make sure we doesn't save unused data

}

// End Update Validation


// Please make sure all data is a correct data before we save it to DB
func (entityStruct *Cart) ValidatebeforeSaveToDB() error {
	// Add you Validation Here

	return nil
}

// End Mutual Function For Create and Update

