package main

import (
	"fmt"
	"os"

	ft "asciiart/features"
)

func main() {
	defaultArgs := ft.CheckSingleArgument(os.Args[1:])
	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		return
	}
	
	colorFlag, outputFlag := ft.CheckColorFlag(defaultArgs), ft.CheckOutputFlag(defaultArgs)
	if colorFlag && outputFlag {
		fmt.Fprintf(os.Stderr, "you can't use '--output' and '--color' in the same commend!\n")
		os.Exit(1)
	}

	flag, args := ft.ExtractFlags(defaultArgs)

	if err := ft.CheckArguments(args); err {
		ft.Usage("color")
		os.Exit(0)
	}
	// Specify the ASCII art banner file to use
	input, banner := ft.GetInputAndBanner(args)
	ft.CheckValidInput(input)
	ft.ProcessInput(input, banner, flag)
}
