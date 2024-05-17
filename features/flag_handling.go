package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	FlagDefs = map[string]string{
		"output": "--output=",
		"color":  "--color=",
	}
	Flag string
)

func ExtractFlags(args []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filteredArgs []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		isFlag := false

		for key, prefix := range FlagDefs {
			if strings.HasPrefix(arg, prefix) {
				flagKey, flagValue := ExtractFlagValue(arg)
				flags[flagKey] = flagValue
				isFlag = true
				Flag = key

				if flagKey == "color" && len(args) > i+2 {
					flags["lettersToBeColored"] = args[i+1]
					i++
				}
				break
			}
		}

		if !isFlag {
			filteredArgs = append(filteredArgs, arg)
		}
	}
	if flags["output"] != "" && flags["color"] != "" {
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--color' in the same commend!\n")
		os.Exit(1)
	}
	fmt.Println(flags)
	fmt.Println(filteredArgs)
	return flags, filteredArgs
}

func ExtractFlagValue(flag string) (string, string) {
	splittedFlag := strings.SplitN(flag, "=", 2)

	if len(splittedFlag) < 2 || len(splittedFlag[1]) == 0 {
		Usage()
		os.Exit(0)
	}

	for key, prefix := range FlagDefs {
		if splittedFlag[0] == prefix[:len(prefix)-1] {
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
