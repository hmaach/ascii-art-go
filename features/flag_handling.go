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
		"align",
	}
	Flag string = "align"
)

// ExtractFlags extracts flags and their values from command-line arguments
func ExtractFlags(args []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filteredArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		isFlag := false

		if strings.HasPrefix(arg, "--") && len(filteredArgs) == 0 {
			flagKey, flagValue, found := findFlagAndExtractValue(arg)
			Flag = strings.Trim(strings.SplitN(arg, "=", 2)[0], "-")
			if found {
				flags[flagKey] = flagValue
				isFlag = true
				if flagKey == "color" {
					_, exists := ColorMap[flagValue]
					if !exists {
						PrintColors()
					}
					if len(args) > i+2 && !strings.HasPrefix(args[i+1], "--") {
						flags["lettersToBeColored"] = args[i+1]
						i++
					}
				}
			} else {
				Usage()
			}
		}

		if !isFlag {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	HandleFlagCombination(flags)
	return flags, filteredArgs
}

// ExtractFlagValue splits a flag into its key and value components
func findFlagAndExtractValue(arg string) (string, string, bool) {
	// Check if the argument contains '='
	splittedFlag := strings.SplitN(arg, "=", 2)
	if len(splittedFlag) < 2 || len(splittedFlag[1]) == 0 {
		Usage()
	}

	flagKey := strings.TrimPrefix(splittedFlag[0], "--")
	flagValue := splittedFlag[1]

	for _, key := range FlagDefs {
		if flagKey == key {
			// Append .txt extension if the flag is 'output' OR does not already have it
			if key == "output" {
				if strings.ContainsAny(flagValue, "/\\") {
					fmt.Printf("Error: you can not use '/' or '\\' in the name of the output file\n")
					os.Exit(1)
				}
				if filepath.Ext(flagValue) != ".txt" {
					flagValue += ".txt"
				}
			}
			return key, strings.ToLower(flagValue), true
		}
	}

	Usage()
	return "", "", false
}
