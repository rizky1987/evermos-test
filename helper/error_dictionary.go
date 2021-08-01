package helper

import "fmt"

func IsRequired(input string) string {
	return fmt.Sprintf("%s is Required", input)
}

func NotFound(input string) string {
	return fmt.Sprintf("%s not found", input)
}

func Mismatch(input1, input2 string) string {
	return fmt.Sprintf("%s and %s Mismatch", input1, input2)
}

func FailedGenerateToken(input string) string {
	return fmt.Sprintf("Failed to Generate Token. Error : %s", input)
}
