package asciiart

import (
	"strings"
)

// Check if there are any non-empty lines in the input lines array
func checkEmptyLines(splittedInput []string) bool {
	for _, line := range splittedInput {
		if line != "" {
			return false
		}
	}
	return true
}

// DrawASCIIArt draws ASCII art and colorizes specific substrings
func DrawASCIIArt(
	charactersMap map[rune][]string,
	splittedInput []string,
	flag map[string]string,
	runesToBeColored []rune,
) []string {
	var result []string
	color := flag["color"]

	// check if the input contain only new lines
	emptyLines := checkEmptyLines(splittedInput)
	if emptyLines {
		newLines := strings.Repeat("\n", len(splittedInput)-1)
		result = append(result, newLines)
		return result
	}

	for inputLineIdx, inputLine := range splittedInput {
		var resultLine strings.Builder
		if inputLine == "" {
			result = append(result, "\n")
			continue
		}

		// Find the starting indices of all occurrences of the substring to be colored
		var substringIndices []int
		substringLen := len(flag["lettersToBeColored"])
		if color != "" && substringLen > 0 {
			substringIndices = FindSubStringIndices(inputLine, flag["lettersToBeColored"])
		}
		// Draw each character of the line in ASCII art format
		for line := 0; line < 8; line++ {
			for charIdx, char := range inputLine {
				// Determine if the current character should be colorized
				shouldColorize := substringLen == 0 || isInRange(charIdx, substringIndices, substringLen)
				if color != "" && shouldColorize && char != 32 {
					resultLine.WriteString(Colorize(charactersMap[char][line], color, inputLineIdx))
				} else {
					if char == 32 && flag["align"] == "justify" {
						resultLine.WriteString("{space}")
					} else {
						resultLine.WriteString(charactersMap[char][line])
					}
				}
			}
			resultLine.WriteString("\n")
		}
		result = append(result, resultLine.String())
	}
	return result
}
