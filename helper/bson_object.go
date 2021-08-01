package helper

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func ChangeBsonObjectIdToString(input bson.ObjectId) string {
	hexaString := fmt.Sprintf("%x", string(input))

	return hexaString
}
