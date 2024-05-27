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

	result := DrawASCIIArt(characterMatrix,
		splittedInput,
		hasNonEmptyLines,
		flag,
		runesToBeColored)
	SaveOrPrintResultToFile(flag["output"], result)
}
