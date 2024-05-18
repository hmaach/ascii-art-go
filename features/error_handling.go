package asciiart

import (
	"fmt"
	"os"
)

func CheckArguments(args []string) bool {
	argCount := len(args)
	if argCount == 2 {
		if args[1] != "shadow" && args[1] != "standard" && args[1] != "thinkertoy" {
			fmt.Printf("invalid banner type '%s'\nAvailable banner types are: 'standard' (default), 'shadow', and 'thinkertoy'\n", args[1])
			os.Exit(0)
		}
	}
	if argCount > 2 {
		return true
	}
	return false
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

func Usage() {
	fmt.Fprintf(os.Stderr, "\n   Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	if Flag == "output" {
		fmt.Fprintf(os.Stderr, "   EX: go run . --output=<fileName.txt> something standard\n\n")
	} else if Flag == "color" {
		fmt.Fprintf(os.Stderr, "   EX: go run . --color=<color> <letters to be colored> \"something\"\n\n")
	} else {
		fmt.Fprintf(os.Stderr, "   EX: go run . --output=<fileName.txt> something standard\n")
		fmt.Fprintf(os.Stderr, "OR\n")
		fmt.Fprintf(os.Stderr, "   EX: go run . --color=<color> <letters to be colored> \"something\" shadow\n\n")
	}
	os.Exit(0)
}
