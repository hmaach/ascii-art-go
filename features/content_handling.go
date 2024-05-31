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

func saveResultToFile(result []string, outputPath string) {
	strResult := strings.Join(result, "")
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
	result := strings.Join(lines, "")
	fmt.Printf("%s", result)
}

func SaveOrPrintResultToFile(result []string, flag map[string]string) {
	if flag["output"] != "" {
		saveResultToFile(result, flag["output"])
	} else {
		printResult(result, flag["align"])
	}
}
