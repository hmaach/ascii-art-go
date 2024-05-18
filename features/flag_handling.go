package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	// FlagDefs defines the command line flags for the application
	FlagDefs = map[string]string{
		"output": "--output=",
		"color":  "--color=",
	}
	Flag string // Current flag being processed
)

func ExtractFlags(args []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filteredArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		isFlag := false

		for key, prefix := range FlagDefs {
			if strings.HasPrefix(arg, prefix) {
				// Extract flag key and value
				flagKey, flagValue := ExtractFlagValue(arg) 
				flags[flagKey] = flagValue
				isFlag = true
				Flag = key

				// If the color flag is found, check for the letters to be colored
				if flagKey == "color" && len(args) > i+2 {
					flags["lettersToBeColored"] = args[i+1]
					i++
				}
				break
			}
		}
		// If the argument is not a flag, add it to the filtered arguments ([string] and [banner])
		if !isFlag {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	// Ensure 'output' and 'color' flags are not used together
	if flags["output"] != "" && flags["color"] != "" {
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--color' in the same commend!\n")
		os.Exit(1)
	}
	return flags, filteredArgs
}

func ExtractFlagValue(flag string) (string, string) {
	splittedFlag := strings.SplitN(flag, "=", 2)

	// Validate the flag format
	if len(splittedFlag) < 2 || len(splittedFlag[1]) == 0 {
		Usage()
		os.Exit(0)
	}

	for key, prefix := range FlagDefs {
		if splittedFlag[0] == prefix[:len(prefix)-1] {
			// Append .txt extension if the flag is 'output' and does not already have it
			if key == "output" && filepath.Ext(splittedFlag[1]) != ".txt" {
				splittedFlag[1] += ".txt"
			}
			return key, strings.ToLower(splittedFlag[1])
		}
	}

	Usage()
	os.Exit(0)
	return "", ""
}
