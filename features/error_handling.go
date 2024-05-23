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
			os.Exit(1)
		}
	}
}

func HandleFlagCombination(flags map[string]string) {
	if flags["output"] != "" && flags["color"] != "" {
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--color' in the same command!\n")
		os.Exit(1)
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	if Flag == "output" {
		fmt.Fprintf(os.Stderr, "EX: go run . --output=<fileName.txt> something standard\n")
	} else if Flag == "color" {
		fmt.Fprintf(os.Stderr, "EX: go run . --color=<color> <letters to be colored> \"something\"\n")
	} else {
		fmt.Fprintf(os.Stderr, "   EX: go run . --output=<fileName.txt> something standard\n")
		fmt.Fprintf(os.Stderr, "OR\n")
		fmt.Fprintf(os.Stderr, "   EX: go run . --color=<color> <letters to be colored> \"something\" shadow\n\n")
	}
	os.Exit(0)
}
