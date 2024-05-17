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

var ColorMap = map[string]string{
	"black":   Black,
	"red":     Red,
	"green":   Green,
	"yellow":  Yellow,
	"blue":    Blue,
	"magenta": Magenta,
	"cyan":    Cyan,
	"white":   White,
}

func Colorize(s, color string) string {
	colorCode, exists := ColorMap[color]
	if !exists {
		PrintColors()
		os.Exit(0)
	}

	return fmt.Sprintf("%s%s%s", colorCode, s, Reset)
}

func PrintColors() {
	fmt.Println("Invalid color. Please choose one of the following colors:")
	for name, code := range ColorMap {
		fmt.Printf("- %s%s%s\n", code, name, Reset)
	}
	os.Exit(1)
}
