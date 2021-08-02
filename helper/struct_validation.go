package helper

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
)


func IsNotFoundErrorValidation(input string) bool {

	return strings.Replace(strings.ToLower(input), " ", "", -1) == "notfound"
}

func ConvertObjectIdToString(input bson.ObjectId) string {
	return fmt.Sprintf("%x", string(input))
}
