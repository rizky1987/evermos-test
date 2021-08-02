package helper

import (
	"encoding/hex"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func ChangeBsonObjectIdToString(input bson.ObjectId) string {
	hexaString := fmt.Sprintf("%x", string(input))

	return hexaString
}

func GenerateBsonObjectId() *bson.ObjectId {

	newBsonObject  := bson.NewObjectId()
	return &newBsonObject
}

func CheckingIsBsonObjectId(input string) bool {

	return bson.IsObjectIdHex(input)
}

func ChangeArrayOfHexaIdToBsonObjectId(inputs []string) []*bson.ObjectId{

	var results []*bson.ObjectId

	for _, input := range inputs {
		isObjectIdHex := bson.IsObjectIdHex(input)
		if isObjectIdHex {
			inputInObjectId := bson.ObjectIdHex(input)
			results = append(results, &inputInObjectId)
		}
	}

	return results
}

func IsObjectIdHexValidation(input string) bool {

	if len(input) != 24 {
		return false
	}

	_, err := hex.DecodeString(input)
	return err == nil
}

func ChangeStringOfObjectIdToBsonObjectId(input string) (*bson.ObjectId, error){

	input = TrimWhiteSpace(input)
	isObjectIdHex := bson.IsObjectIdHex(input)
	if isObjectIdHex {
		result := bson.ObjectIdHex(input)

		return &result, nil
	}

	return nil, errors.New(ErrorIsNOTObjectIdHex(input))
}