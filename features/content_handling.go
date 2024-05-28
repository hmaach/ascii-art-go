package asciiart

import (
	"fmt"
	"strings"
)

// Convert content array to a character matrix mapping ASCII characters to their line representations
func ConvertToCharacterMatrix(content []string) map[rune][]string {
	characterMatrix := map[rune][]string{}
	for i, val := range content {
		characterMatrix[rune(32+i)] = strings.Split(val, "\n")
	}
	return characterMatrix
}

// Check if there are any non-empty lines in the input lines array
func CheckEmptyLines(splittedInput []string) bool {
	for _, line := range splittedInput {
		if line != "" {
			return true
		}
	}
	return false
}

func SaveOrPrintResultToFile(result string, flag map[string]string) {
	if flag["output"] != "" {
		err := SaveFile(flag["output"], result)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		if flag["align"] != "" && flag["align"] != "left" {
			result = Justify(result, flag["align"])
		}
		fmt.Printf("%s", result)
	}
}
