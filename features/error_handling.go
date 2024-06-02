package asciiart

import (
	"errors"
	"fmt"
	"os"
)

func CheckArguments(args []string) error {
	argCount := len(args)
	if argCount == 2 {
		switch args[1] {
		case "shadow", "standard", "thinkertoy":
			// valid banner types, do nothing
		default:
			return fmt.Errorf("invalid banner type '%s'\nAvailable banner types are: 'standard' (default), 'shadow', and 'thinkertoy'", args[1])
		}
	} else if argCount > 2 {
		return errors.New("too many arguments")
	}
	return nil
}

// validates if the input contains only printable ASCII characters
func CheckValidInput(input string) {
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			fmt.Println("Error: The input contains characters without corresponding ASCII art representation!")
			os.Exit(2)
		}
	}
}

func HandleFlagCombination(flags map[string]string) {
	switch {
	case flags["output"] != "" && flags["color"] != "":
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--color' in the same command!\n")
		os.Exit(1)
	case flags["output"] != "" && flags["align"] != "":
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--align' in the same command!\n")
		os.Exit(2)
	}
}

func PrintColors() {
	fmt.Println("Invalid color. Please choose one of the following colors:")
	for name, code := range ColorMap {
		fmt.Printf("- %s%s%s\n", code, name, Reset)
	}
	os.Exit(2)
}

func PrintAlignments() {
	fmt.Println("Invalid alignment. Please choose one of the following alignment:")
	for _, align := range Alignments {
		fmt.Printf("- %s\n", align)
	}
	os.Exit(2)
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	switch Flag {
	case "output":
		fmt.Fprintf(os.Stderr, "EX: go run . --output=<fileName.txt> something standard\n")
	case "color":
		fmt.Fprintf(os.Stderr, "EX: go run . --color=<color> <letters to be colored> \"something\"\n")
	case "align":
		fmt.Fprintf(os.Stderr, "Example: go run . --align=right \"something\" standard\n")
	default:
		fmt.Fprintf(os.Stderr, "   EX: go run . --output=<fileName.txt> something standard\n")
		fmt.Fprintf(os.Stderr, "OR\n")
		fmt.Fprintf(os.Stderr, "   EX: go run . --color=<color> <letters to be colored> \"something\" shadow\n\n")
	}
	os.Exit(2)
}
