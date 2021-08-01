package helper

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func IsObjectIdHexValidation(input string) bool {

	if len(input) != 24 {
		return false
	}

	_, err := hex.DecodeString(input)
	return err == nil
}

func IsNotFoundErrorValidation(input string) bool {

	return strings.Replace(strings.ToLower(input), " ", "", -1) == "notfound"
}

func ConvertObjectIdToString(input bson.ObjectId) string {
	return fmt.Sprintf("%x", string(input))
}
