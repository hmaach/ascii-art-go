package main

import (
	"fmt"
	"os"
	"strings"

	ft "asciiart/features"
)

func main() {
	defaultArgs := os.Args[1:]
	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		return
	}

	outputFile, args := ft.ExtractOutputFlag(defaultArgs)

	if err := ft.CheckArguments(args); err {
		// fmt.Println(err)
		ft.Usage()
		os.Exit(0)
	}

	if outputFile != "" {
		if err := ft.ValidateFileExtension(outputFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(0)
		}
	}

	// Specify the ASCII art banner file to use
	var banner string
	if len(args) == 2 {
		banner = args[1]
		banner += ".txt"
	} else {
		banner = "standard.txt"
	}

	var input string
	if len(args) >= 1 {
		input = args[0]
	} else {
		input = ""
	}

	if !ft.CheckValidInput(input) {
		fmt.Println("Error: The input contains characters without corresponding ASCII art representation!")
		os.Exit(1)
	}

	data, err := os.ReadFile("banners/" + banner)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	splittedInput := strings.Split(input, "\\n")
	hasNonEmptyLines := ft.CheckEmptyLines(splittedInput)

	stringData := string(data[1:])
	if banner == "thinkertoy.txt" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}
	content := strings.Split(stringData, "\n\n")
	characterMatrix := ft.ConvertToCharacterMatrix(content)

	result := ft.DrawASCIIArt(characterMatrix, splittedInput, hasNonEmptyLines)

	if outputFile != "" {
		err := ft.SaveFile(outputFile, result)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Printf("%s", result)
	}
}
