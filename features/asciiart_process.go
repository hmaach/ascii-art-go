package asciiart

import (
	"strings"
)

func handleSpaces(splittedInput []string) []string {
	for i, line := range splittedInput {
		tempLine := strings.Fields(line)
		splittedInput[i] = strings.Join(tempLine, " ")
	}
	return splittedInput
}

// ProcessInput processes the input string, reads the banner, and produces the ASCII art
func ProcessInput(input, banner string, flags map[string]string) {
	splittedInput := strings.Split(input, "\\n")

	if flags["color"] != "" {
		for range splittedInput {
			SpacesOfColor = append(SpacesOfColor, 0)
		}
	}

	if flags["align"] == "justify" {
		splittedInput = handleSpaces(splittedInput)
	}

	charactersMap := ReadBanner(banner)

	var runesToBeColored []rune
	lettersToBeColored := flags["lettersToBeColored"]

	if lettersToBeColored != "" && strings.Contains(input, lettersToBeColored) {
		runesToBeColored = []rune(lettersToBeColored)
	}

	result := DrawASCIIArt(charactersMap,
		splittedInput,
		flags,
		runesToBeColored)
	SaveOrPrintResultToFile(result, flags)
}
