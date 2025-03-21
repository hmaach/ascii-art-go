package asciiart

import (
	"fmt"
	"os"
	"strings"
)

// Convert content array to a character mapping ASCII characters to their line representations
func ConvertTocharacterMap(content []string) map[rune][]string {
	charactersMap := map[rune][]string{}
	for i, val := range content {
		charactersMap[rune(32+i)] = strings.Split(val, "\n")
	}
	return charactersMap
}

func saveResultToFile(result []string, outputPath string) {
	strResult := strings.Join(result, "")
	err := SaveFile(outputPath, strResult)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// printResult processes and prints lines based on the specified alignment.
func printResult(lines []string, flags map[string]string) {
	alignment := flags["align"]

	if alignment != "" && alignment != "left" {
		for i, line := range lines {
			justifiedLine, err := Justify(line, i, flags)
			if err != nil {
				fmt.Printf("Error justifying line: %v\n", err)
				os.Exit(2)
			}
			lines[i] = justifiedLine
		}
	}
	result := strings.Join(lines, "")
	fmt.Printf("%s", result)
}

func SaveOrPrintResultToFile(result []string, flags map[string]string) {
	if flags["output"] != "" {
		saveResultToFile(result, flags["output"])
	} else {
		printResult(result, flags)
	}
}
