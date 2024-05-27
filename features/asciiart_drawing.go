package asciiart

import "strings"

// DrawASCIIArt draws ASCII art and colorizes specific substrings
func DrawASCIIArt(
	characterMatrix map[rune][]string,
	splittedInput []string,
	hasNonEmptyLines bool,
	flag map[string]string,
	runesToBeColored []rune) string {

	var result strings.Builder
	color := flag["color"]

	for i, val := range splittedInput {
		if val == "" {
			if hasNonEmptyLines || i != 0 {
				result.WriteString("\n")
			}
			continue
		}

		// Find the starting indices of all occurrences of the substring to be colored
		substringIndices := FindSubStringIndices(val, flag["lettersToBeColored"])
		substringLen := len(flag["lettersToBeColored"])

		// Draw each character of the line in ASCII art format
		for j := 0; j < 8; j++ {
			for kIdx, k := range val {
				// Determine if the current character should be colorized
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
