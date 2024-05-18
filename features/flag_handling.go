package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	// FlagDefs defines the command line flags for the application
	FlagDefs = []string{
		"output",
		"color",
	}
	Flag string // Current flag being processed
)

func ExtractFlags(args []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filteredArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		isFlag := false
		if strings.HasPrefix(arg, "--") {
			Flag = strings.Trim(arg,"-=") // use Flag to display the Usage based on the flag
			for _, key := range FlagDefs {
				if strings.Contains(arg, key) {
					// Extract flag key and value
					flagKey, flagValue := ExtractFlagValue(arg)
					flags[flagKey] = flagValue
					isFlag = true
					// If the color flag is found, check for the letters to be colored
					if flagKey == "color" && len(args) > i+1 {
						flags["lettersToBeColored"] = args[i+1]
						i++
					}
					break
				}
			}
		}
		// If the argument is not a flag, add it to the filtered arguments ([string] and [banner])
		if !isFlag {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	// Ensure 'output' and 'color' flags are not used together
	if flags["output"] != "" && flags["color"] != "" {
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--color' in the same command!\n")
		os.Exit(1)
	}
	return flags, filteredArgs
}

// ExtractFlagValue splits a flag into its key and value components
func ExtractFlagValue(flag string) (string, string) {
	splittedFlag := strings.SplitN(flag, "=", 2)

	// Handle missing value case
	if len(splittedFlag) < 2 || len(splittedFlag[1]) == 0 {
		Usage()
		os.Exit(0)
	}

	for _, prefix := range FlagDefs {
		if splittedFlag[0] == "--"+prefix {
			// Append .txt extension if the flag is 'output' and does not already have it
			if prefix == "output" && filepath.Ext(splittedFlag[1]) != ".txt" {
				splittedFlag[1] += ".txt"
			}
			return prefix, strings.ToLower(splittedFlag[1])
		}
	}

	Usage()
	os.Exit(0)
	return "", ""
}
