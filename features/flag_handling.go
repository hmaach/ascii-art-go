package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func ExtractOutputFlag(args []string) (map[string]string, []string) {
	flags := make(map[string]string)
	var filteredArgs []string

	for i, arg := range args {
		if strings.HasPrefix(arg, "--output=") && i == 0 {
			flags["outputFile"] = extractOutputFileName(arg)
		} else if strings.HasPrefix(arg, "--color=") && i == 0 {
			flags["color"] = extractColor(arg)
			if flags["color"] != "" && len(args) == 4 {
				flags["lettersToBeColored"] = args[i+1]
			}
		} else if strings.HasPrefix(arg, "--") {
			Usage()
			os.Exit(0)
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	return flags, filteredArgs
}

func extractOutputFileName(flag string) string {
	fileName := strings.TrimPrefix(flag, "--output=")
	if len(fileName) == 0 {
		Usage()
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
		Usage()
		os.Exit(0)
	}
	return color
}

// func ExtractOutputFlag(args []string) (string, []string) {
// 	var outputFile string
// 	var filteredArgs []string

// 	// check if arguments are given as one string
// 	if len(args) == 1 {
// 		tempArgs := strings.Fields(args[0])
// 		for _, v := range tempArgs {
// 			if strings.HasPrefix(v, "--") {
// 				args = tempArgs
// 			}
// 		}
// 	}

// 	for i, arg := range args {
// 		if strings.HasPrefix(arg, "--output=") && i == 0 {
// 			outputFile = strings.TrimPrefix(arg, "--output=")
// 			if len(outputFile) == 0 {
// 				// fmt.Printf("you must include the output file!")
// 				Usage()
// 				os.Exit(0)
// 			}
// 		} else if strings.HasPrefix(arg, "--") {
// 			Usage()
// 			os.Exit(0)
// 		} else {
// 			filteredArgs = append(filteredArgs, arg)
// 		}
// 	}
// 	if outputFile != "" {
// 		if err := ValidateFileExtension(outputFile); err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			os.Exit(0)
// 		}
// 	}
// 	return outputFile, filteredArgs
// }
