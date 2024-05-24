package main

import (
	"fmt"
	"os"

	ft "asciiart/features"
)

func main() {
	defaultArgs := os.Args[1:]
	if len(defaultArgs) == 0 || len(defaultArgs[0]) == 0 {
		if len(defaultArgs) > 1 && len(defaultArgs[0]) == 0 {
			ft.Usage()
		}
		return
	}

	flag, args := ft.ExtractFlags(defaultArgs)

	if err := ft.CheckArguments(args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Specify the ASCII art banner file to use
	input, banner := ft.GetInputAndBanner(args)
	ft.CheckValidInput(input)
	ft.ProcessInput(input, banner, flag)
}
