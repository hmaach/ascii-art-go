package asciiart

import (
	"fmt"
	"strings"
)

// ProcessInput processes the input string, reads the banner, and produces the ASCII art
func ProcessInput(input, banner string, flag map[string]string) {
	splittedInput := strings.Split(input, "\\n")
	hasNonEmptyLines := CheckEmptyLines(splittedInput)
	characterMatrix := ReadBanner(banner)

	var runesToBeColored []rune
	lettersToBeColored := flag["lettersToBeColored"]

	if lettersToBeColored != "" && strings.Contains(input, lettersToBeColored) {
		runesToBeColored = []rune(lettersToBeColored)
	}

	width, errWidth := GetTerminalWidth()
	if errWidth != nil {
		fmt.Printf("Error getting terminal width: %v\n", errWidth)
	}
	fmt.Println(width)

	result := DrawASCIIArt(characterMatrix, splittedInput, hasNonEmptyLines, flag, runesToBeColored)
	SaveOrPrintResultToFile(flag["output"], result)
}

// DrawASCIIArt draws ASCII art and colorizes specific substrings
func DrawASCIIArt(characterMatrix map[rune][]string, splittedInput []string, hasNonEmptyLines bool, flag map[string]string, runesToBeColored []rune) string {
	var result strings.Builder
	color := flag["color"]

	for i, val := range splittedInput {
		if val == "" {
			if hasNonEmptyLines || i != 0 {
				result.WriteString("\n")
			}
			continue
		}

		substringIndices := FindSubStringIndices(val, flag["lettersToBeColored"])
		substringLen := len(flag["lettersToBeColored"])

		for j := 0; j < 8; j++ {
			for kIdx, k := range val {
				shouldColorize := substringLen == 0 || isInRange(kIdx, substringIndices, substringLen)
				if color != "" && shouldColorize {
					result.WriteString(Colorize(characterMatrix[k][j], color))
				} else {
					result.WriteString(characterMatrix[k][j])
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}

// Checks if the given index is within any of the ranges specified by the start indices and length
func isInRange(index int, indices []int, length int) bool {
	for _, start := range indices {
		if index >= start && index < start+length {
			return true
		}
	}
	return false
}

// Returns the start indices of all occurrences of the substring in the main string
func FindSubStringIndices(str, substr string) []int {
	var indices []int
	if len(substr) > 0 {
		for i := 0; ; {
			index := strings.Index(str[i:], substr)
			if index == -1 {
				break
			}
			indices = append(indices, i+index)
			i += index + len(substr)
		}
	}
	return indices
}
