package asciiart

import (
	"fmt"
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
	Orange  = "\033[0;38;5;208m"
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
	"orange":  Orange,
	"white":   White,
}

func Colorize(s, color string) string {
	colorCode, exists := ColorMap[color]
	if !exists {
		PrintColors()
	}

	return fmt.Sprintf("%s%s%s", colorCode, s, Reset)
}
