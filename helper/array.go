package helper

import (
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
		if errorObjectId != nil {

			errorResults = append(errorResults, errorObjectId.Error())
		} else {

			results = append(results, objectId)
		}
	}

	return results, errorResults
}

func RemoveDuplicateArrayTrimWhiteSpace(inputStringArrayStatus []string) []string {

	results := []string{}

	if len(inputStringArrayStatus) > 0 {

		results = RemoveDuplicateArrayAndTrimWhiteSpace(inputStringArrayStatus)
	}

	return results
}
