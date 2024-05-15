package main

import (
	"os"

	ft "asciiart/features"
)

func main() {
	defaultArgs := ft.CheckSingleArgument(os.Args[1:])
	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		return
	}
	flag, args := ft.ExtractOutputFlag(defaultArgs)
	if err := ft.CheckArguments(args); err {
		ft.Usage()
		os.Exit(0)
	}
	// Specify the ASCII art banner file to use
	input, banner := ft.GetInputAndBanner(args)
	ft.CheckValidInput(input)
	ft.ProcessInput(input, banner, flag)
}
