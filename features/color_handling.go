package asciiart

import (
	"fmt"
	"os"
)

const (
	Black   = "\033[0;30m"
	Red     = "\033[0;31m"
	Green   = "\033[0;32m"
	Yellow  = "\033[0;33m"
	Blue    = "\033[0;34m"
	Magenta = "\033[0;35m"
	Cyan    = "\033[0;36m"
	White   = "\033[0;37m"
	Reset   = "\033[0m"
)

func AreLettersToBeColored(chars, str string) bool {
	if len(chars) > len(str) {
		return false
	}
	for _, char := range chars {
		found := false
		for _, s := range str {
			if char == s {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func ConvertLettersToSliceOfRunes(str string) []rune {
	result := []rune{}
	for _, char := range str {
		result = append(result, char)
	}
	return result
}

func PrintStringInColor(s, color string) {
	var colorCode string

	switch color {
	case "black":
		colorCode = Black
	case "red":
		colorCode = Red
	case "green":
		colorCode = Green
	case "yellow":
		colorCode = Yellow
	case "blue":
		colorCode = Blue
	case "magenta":
		colorCode = Magenta
	case "cyan":
		colorCode = Cyan
	case "white":
		colorCode = White
	default:
		PrintColors()
	}

	fmt.Printf("%s%s%s", colorCode, s, Reset)
}

func IsInSlice(k rune, slice []rune) bool {
	for _, v := range slice {
		if v == k {
			return true
		}
	}
	return false
}

func PrintColors() {
	fmt.Println("Invalid color. Please choose one of the following colors:")
	fmt.Printf("- %sblack%s\n", Black, Reset)
	fmt.Printf("- %sred%s\n", Red, Reset)
	fmt.Printf("- %sgreen%s\n", Green, Reset)
	fmt.Printf("- %syellow%s\n", Yellow, Reset)
	fmt.Printf("- %sblue%s\n", Blue, Reset)
	fmt.Printf("- %smagenta%s\n", Magenta, Reset)
	fmt.Printf("- %scyan%s\n", Cyan, Reset)
	fmt.Printf("- %swhite%s\n", White, Reset)
	os.Exit(0)
}
