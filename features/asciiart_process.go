package asciiart

import (
	"strings"
)

// ProcessInput processes the input string, reads the banner, and produces the ASCII art
func ProcessInput(input, banner string, flags map[string]string) {
	splittedInput := strings.Split(input, "\\n")
	hasNonEmptyLines := CheckEmptyLines(splittedInput)
	characterMatrix := ReadBanner(banner)

	var runesToBeColored []rune
	lettersToBeColored := flags["lettersToBeColored"]

	if lettersToBeColored != "" && strings.Contains(input, lettersToBeColored) {
		runesToBeColored = []rune(lettersToBeColored)
	}

	result := DrawASCIIArt(characterMatrix,
		splittedInput,
		hasNonEmptyLines,
		flags,
		runesToBeColored)
	SaveOrPrintResultToFile(result, flags)
}
