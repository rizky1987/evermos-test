package helper

import "fmt"

func ErrorIsRequired(input string) string {
	return fmt.Sprintf("%s is Required", input)
}

func ErrorNotFound(input string) string {
	return fmt.Sprintf("%s not found", input)
}

func ErrorMismatch(input1, input2 string) string {
	return fmt.Sprintf("%s and %s Mismatch", input1, input2)
}

func ErrorFailedGenerateToken(input string) string {
	return fmt.Sprintf("Failed to Generate Token. Error : %s", input)
}

func ErrorIsNOTObjectIdHex(input string) string {
	return fmt.Sprintf("Your Inputed ID [%s] is not a hexa string", input)
}
