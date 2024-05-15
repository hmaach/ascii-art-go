package asciiart

import (
	"fmt"
	"strings"
)

func ProcessInput(input, banner string, flag map[string]string) {
	splittedInput := strings.Split(input, "\\n")
	hasNonEmptyLines := CheckEmptyLines(splittedInput)
	data := ReadBanner(banner)

	if banner == "thinkertoy.txt" {
		data = strings.ReplaceAll(data, "\r", "")
	}
	content := strings.Split(data, "\n\n")
	characterMatrix := ConvertToCharacterMatrix(content)

	if flag["color"] != "" {
		DrawASCIIArt2(characterMatrix, splittedInput, hasNonEmptyLines, flag)
	} else {
		result := DrawASCIIArt(characterMatrix, splittedInput, hasNonEmptyLines)
		SaveOrPrintResultToFile(flag["outputFile"], result)
	}
}

// Render the ASCII art based on the character matrix and the input lines
func DrawASCIIArt(characterMatrix map[rune][]string, splittedInput []string, hasNonEmptyLines bool) string {
	result := ""
	for i, val := range splittedInput {
		if val == "" {
			if hasNonEmptyLines {
				result += "\n"
			} else if i != 0 && !hasNonEmptyLines {
				result += "\n"
			}
		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					result += characterMatrix[k][j]
				}
				result += "\n"
			}
		}
	}
	return result
}

// Render the ASCII art based on the character matrix and the input lines
func DrawASCIIArt2(characterMatrix map[rune][]string, splittedInput []string, hasNonEmptyLines bool, color map[string]string) string {
	result := ""
	for i, val := range splittedInput {
		if val == "" {
			if hasNonEmptyLines {
				fmt.Println()
			} else if i != 0 && !hasNonEmptyLines {
				fmt.Println()
			}
		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					fmt.Printf(characterMatrix[k][j])
				}
				fmt.Println()
			}
		}
	}
	return result
}
