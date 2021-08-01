package api_helper

import (
	"evermos-test/helper"
	"fmt"
)

func GenerateUsername(homeBlockId, homeNumberId string) string {

	username := fmt.Sprintf("%v-%v",
		helper.TrimWhiteSpace(homeBlockId),
		helper.TrimWhiteSpace(homeNumberId))

	return username
}
