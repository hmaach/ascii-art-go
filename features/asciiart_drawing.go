package asciiart

import "strings"

// DrawASCIIArt draws ASCII art and colorizes specific substrings
func DrawASCIIArt(
	characterMatrix map[rune][]string,
	splittedInput []string,
	hasNonEmptyLines bool,
	flag map[string]string,
	runesToBeColored []rune,
) string {
	var result strings.Builder
	color := flag["color"]

	for i, inputLine := range splittedInput {
		if inputLine == "" {
			if hasNonEmptyLines || i != 0 {
				result.WriteString("\n")
			}
			continue
		}

		// Find the starting indices of all occurrences of the substring to be colored
		substringIndices := FindSubStringIndices(inputLine, flag["lettersToBeColored"])
		substringLen := len(flag["lettersToBeColored"])

		// Draw each character of the line in ASCII art format
		for line := 0; line < 8; line++ {
			for charIdx, char := range inputLine {
				// Determine if the current character should be colorized
				shouldColorize := substringLen == 0 || isInRange(charIdx, substringIndices, substringLen)
				if color != "" && shouldColorize && char != 32 {
					result.WriteString(Colorize(characterMatrix[char][line], color))
				} else {
					if char == 32 && flag["align"] == "justify" {
						result.WriteString("{space}")
					} else {
						result.WriteString(characterMatrix[char][line])
					}
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}
