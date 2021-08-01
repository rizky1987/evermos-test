package helper

import (
	"regexp"
	"strconv"
	"strings"
)

func ChangeStringToInt64(inputString string) (int64, error) {
	int64Result, err := strconv.ParseInt(TrimWhiteSpace(inputString), 10, 64)
	if err != nil {
		return 0, err
	}
	return int64Result, nil
}

func ChangeStringToInt(inputString string) (int, error) {
	intResult, err := strconv.Atoi(TrimWhiteSpace(inputString))
	if err != nil {
		return 0, err
	}
	return intResult, nil
}

func ChangeStringToBool(inputString string) (bool, error) {

	stringTrim := TrimWhiteSpace(inputString)
	if stringTrim == "" {
		return false, nil
	}
	isTrueOrFalse, err := strconv.ParseBool(stringTrim)
	if err != nil {
		return false, err
	}
	return isTrueOrFalse, nil
}

func ChangeStringToFloat64(inputString string) (float64, error) {

	float64Value, err := strconv.ParseFloat(TrimWhiteSpace(inputString), 64)
	if err != nil {
		return 0, err
	}

	return float64Value, nil
}

func TrimWhiteSpace(inputString string) string {
	return strings.TrimSpace(inputString)
}

func TrimWhiteSpaceArrayString(inputStrings []string) []string {

	results := []string{}

	for i := 0; i < len(inputStrings); i++ {
		results = append(results, TrimWhiteSpace(inputStrings[i]))
	}

	return results
}

func SplitStringToArrayString(inputString, delimiter string) []string {
	return strings.Split(TrimWhiteSpace(inputString), TrimWhiteSpace(delimiter))
}

func CountDuplicateValueInStringOfArray(inputStringArray []string, delimiter string) int {
	count := 0

	for _, str := range inputStringArray {
		if TrimWhiteSpace(str) == TrimWhiteSpace(delimiter) {
			count++
		}
	}

	return count
}

func RemoveDuplicatesStringArray(inputStringArrayArticleIds, inputStringArrayStatus []string) []string {

	keys := make(map[string]bool)
	listResult := []string{}
	for i, entry := range inputStringArrayArticleIds {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			listResult = append(listResult, entry)
		} else {
			inputStringArrayStatus[i] = inputStringArrayStatus[len(inputStringArrayStatus)-1] // Copy last element to index i.
			inputStringArrayStatus[len(inputStringArrayStatus)-1] = ""                        // Erase last element (write zero value).
			inputStringArrayStatus = inputStringArrayStatus[:len(inputStringArrayStatus)-1]   // Truncate slice.
		}
	}
	return listResult
}

func FindInArrayString(value string, arrayString []string) bool {
	for _, entry := range arrayString {
		if value == entry {
			return true
		}
	}
	return false
}

func CheckStringRegex(inputString string) bool {

	usernameRegex := regexp.MustCompile("^[a-zA-Z0-9]*$")
	if !usernameRegex.MatchString(inputString) {
		return false
	}

	return true
}
