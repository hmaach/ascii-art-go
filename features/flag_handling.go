package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func CheckOutputFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			return true
		}
	}
	return false
}

func CheckColorFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			return true
		}
	}
	return false
}

func ExtractFlags(args []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filteredArgs []string
	var indexOfToBeColored int

	for i, arg := range args {
		if strings.HasPrefix(arg, "--output=") && i == 0 {
			flags["outputFile"] = extractOutputFileName(arg)
		} else if strings.HasPrefix(arg, "--color=") && i == 0 {
			flags["color"] = extractColor(arg)
			if flags["color"] != "" && i+2 < len(args) {
				if AreLettersToBeColored(args[i+1], args[i+2]) {
					flags["lettersToBeColored"] = args[i+1]
					indexOfToBeColored = i + 1
				}
			}
		} else if strings.HasPrefix(arg, "--") {
			Usage("color")
			os.Exit(0)
		} else {
			if i != indexOfToBeColored {
				filteredArgs = append(filteredArgs, arg)
			}
		}
	}
	return flags, filteredArgs
}

func extractOutputFileName(flag string) string {
	fileName := strings.TrimPrefix(flag, "--output=")
	if len(fileName) == 0 {
		Usage("output")
		os.Exit(0)
	}
	if fileName != "" {
		if err := ValidateFileExtension(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
	}
	return fileName
}

func extractColor(flag string) string {
	color := strings.TrimPrefix(flag, "--color=")
	if len(color) == 0 {
		Usage("color")
		os.Exit(0)
	}
	return color
}
