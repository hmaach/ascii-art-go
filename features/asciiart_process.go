package asciiart

import (
	"slices"
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

	var runesToBeColored []rune
	colorizeAll := true
	if flag["lettersToBeColored"] != "" {
		if strings.Contains(input, flag["lettersToBeColored"]) {
			runesToBeColored = []rune(flag["lettersToBeColored"])
			colorizeAll = false
		} else {
		}
	}

	result := DrawASCIIArt(characterMatrix, splittedInput, hasNonEmptyLines, colorizeAll, flag, runesToBeColored)
	SaveOrPrintResultToFile(flag["output"], result)
}

func DrawASCIIArt(characterMatrix map[rune][]string, splittedInput []string, hasNonEmptyLines, colorizeAll bool, flag map[string]string, runesToBeColored []rune) string {
	result := ""
	for i, val := range splittedInput {
		if val == "" {
			if hasNonEmptyLines {
				result += "\n"
			} else if i != 0 && !hasNonEmptyLines {
				result += "\n"
			}
		} else {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					if flag["color"] != "" {
						if len(runesToBeColored) > 0 && slices.Contains(runesToBeColored, k) {
							result += Colorize(characterMatrix[k][j], flag["color"])
						} else if len(runesToBeColored) == 0 && !colorizeAll {
							result += Colorize(characterMatrix[k][j], flag["color"])
						} else {
							result += characterMatrix[k][j]
						}
					} else {
						result += characterMatrix[k][j]
					}
				}
				result += "\n"
			}
		}
	}
	return result
}
