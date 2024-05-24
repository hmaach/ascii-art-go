package asciiart

import (
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
			if found {
				Flag = strings.Trim(arg, "-=")
				flags[flagKey] = flagValue
				isFlag = true
				if flagKey == "color" && len(args) > i+2 {
					flags["lettersToBeColored"] = args[i+1]
					i++
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

// findFlagAndExtractValue searches for a flag in a given argument and extracts its value
func findFlagAndExtractValue(arg string) (string, string, bool) {
	for _, key := range FlagDefs {
		if strings.Contains(arg, key) {
			flagKey, flagValue := extractFlagValue(arg)
			return flagKey, flagValue, true
		}
	}
	return "", "", false
}

// ExtractFlagValue splits a flag into its key and value components
func extractFlagValue(flag string) (string, string) {
	splittedFlag := strings.SplitN(flag, "=", 2)

	// Handle missing value case
	if len(splittedFlag) < 2 || len(splittedFlag[1]) == 0 {
		Usage()
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
	return "", ""
}
