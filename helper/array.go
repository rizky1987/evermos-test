package helper

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func RemoveDuplicateArrayAndTrimWhiteSpace(inputStringArrayStatus []string) []string {

	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v:= range inputStringArrayStatus {
		encountered[inputStringArrayStatus[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, TrimWhiteSpace(key))
	}

	return result
}

func RemoveDuplicateArrayOfStringOfBsonIdAndChangeToBsonId(inputStringArrayStatus []string) ([]*bson.ObjectId, []string) {

	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v:= range inputStringArrayStatus {
		encountered[inputStringArrayStatus[v]] = true
	}

	// Place all keys from the map into a slice.
	var results []*bson.ObjectId
	var errorResults []string
	for key, _ := range encountered {

		objectId, errorObjectId :=ChangeStringOfObjectIdToBsonObjectId(key)
		if errorObjectId != "" {

			errorResults = append(errorResults, errorObjectId)
		} else {

			results = append(results, objectId)
		}
	}

	return results, errorResults
}

func ChangeStringOfObjectIdToBsonObjectId(input string) (*bson.ObjectId, string){

	input = TrimWhiteSpace(input)
	isObjectIdHex := bson.IsObjectIdHex(input)
	if isObjectIdHex {
		result := bson.ObjectIdHex(input)

		return &result, ""
	}

	errMessage := fmt.Sprintf("Cannot convert %v to bson Id", input)
	return nil, errMessage
}

func RemoveDuplicateArrayTrimWhiteSpace(inputStringArrayStatus []string) []string {

	results := []string{}

	if len(inputStringArrayStatus) > 0 {

		results = RemoveDuplicateArrayAndTrimWhiteSpace(inputStringArrayStatus)
	}

	return results
}
