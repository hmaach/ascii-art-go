package asciiart

import "strings"

// DrawASCIIArt draws ASCII art and colorizes specific substrings
func DrawASCIIArt(
	characterMatrix map[rune][]string,
	splittedInput []string,
	hasNonEmptyLines bool,
	flag map[string]string,
	runesToBeColored []rune,
) []string {
	var result []string
	color := flag["color"]
	for i, inputLine := range splittedInput {
		var resultLine strings.Builder
		if inputLine == "" {
			if hasNonEmptyLines || i != 0 {
				resultLine.WriteString("\n")
			}
			continue
		}
		// Find the starting indices of all occurrences of the substring to be colored
		var substringIndices []int
		var substringLen int
		if color != "" {
			substringIndices = FindSubStringIndices(inputLine, flag["lettersToBeColored"])
			substringLen = len(flag["lettersToBeColored"])
		}
		// Draw each character of the line in ASCII art format
		for line := 0; line < 8; line++ {
			for charIdx, char := range inputLine {
				// Determine if the current character should be colorized
				shouldColorize := substringLen == 0 || isInRange(charIdx, substringIndices, substringLen)
				if color != "" && shouldColorize && char != 32 {
					resultLine.WriteString(Colorize(characterMatrix[char][line], color))
				} else {
					if char == 32 && flag["align"] == "justify" {
						resultLine.WriteString("{space}")
					} else {
						resultLine.WriteString(characterMatrix[char][line])
					}
				}
			}
			resultLine.WriteString("\n")
		}
		result = append(result, resultLine.String())
	}
	return result
}
