package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExtractOutputFlag(args []string) (string, []string) {
	var outputFile string
	var filteredArgs []string

	// check if arguments are given as one string
	if len(args) == 1 {
		tempArgs := strings.Fields(args[0])
		for _, v := range tempArgs {
			if strings.HasPrefix(v, "--") {
				args = tempArgs
			}
		}
	}

	for i, arg := range args {
		if strings.HasPrefix(arg, "--output=") && i == 0 {
			outputFile = strings.TrimPrefix(arg, "--output=")
			if len(outputFile) == 0 {
				// fmt.Printf("you must include the output file!")
				Usage()
				os.Exit(0)
			}
		} else if strings.HasPrefix(arg, "--") {
			// If arg starts with "--" and not "--output=", it's an invalid flag
			// fmt.Printf("Error: Invalid flag \"%s\". Flag must be in the format '--output=<fileName.txt>'.", arg)
			Usage()
			os.Exit(0)
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	return outputFile, filteredArgs
}

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
func CheckValidInput(input string) bool {
	for _, char := range input {
		if int(char) < 32 || int(char) > 126 {
			return false
		}
	}
	return true
}

func ValidateFileExtension(filename string) error {
	acceptableExtensions := []string{".txt", ".json"}
	extension := strings.ToLower(filepath.Ext(filename))
	if extension == "" {
		return fmt.Errorf("please use one of the following extensions for the output file: .txt")
	}
	for _, ext := range acceptableExtensions {
		if extension == ext {
			return nil
		}
	}
	return fmt.Errorf("invalid file extension '%s' for --output option. Please use one of the following extensions: .txt", extension)
}

func Usage() {
	fmt.Fprintf(os.Stderr, "\n   Usage: go run . [OPTION] [STRING] [BANNER]\n\n   Example: go run . --output=<fileName.txt> something standard\n\n")
}
