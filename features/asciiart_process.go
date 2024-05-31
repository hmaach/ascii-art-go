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

	if flags["align"] == "justify" {
		splittedInput = handleSpaces(splittedInput)
	}

	characterMatrix := ReadBanner(banner)

	var runesToBeColored []rune
	lettersToBeColored := flags["lettersToBeColored"]

	if lettersToBeColored != "" && strings.Contains(input, lettersToBeColored) {
		runesToBeColored = []rune(lettersToBeColored)
	}

	result := DrawASCIIArt(characterMatrix,
		splittedInput,
		flags,
		runesToBeColored)
	SaveOrPrintResultToFile(result, flags)
}
