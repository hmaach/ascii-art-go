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

func saveResultToFile(result []string, outputPath string) {
	strResult := strings.Join(result, "\n")
	err := SaveFile(outputPath, strResult)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func printResult(lines []string, alignment string) {
	if alignment != "" && alignment != "left" {
		for i, line := range lines {
			lines[i] = Justify(line, alignment)
		}
	}
	result := strings.Join(lines, "\n")
	fmt.Printf("%s", result)
}

func SaveOrPrintResultToFile(result []string, flag map[string]string) {
	if flag["output"] != "" {
		saveResultToFile(result, flag["output"])
	} else {
		printResult(result, flag["align"])
	}
}
