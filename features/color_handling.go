package asciiart

import (
	"fmt"
	"strings"
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

var SpacesOfColor []int

func Colorize(s, color string, inputLineIdx int) string {
	colorCode := ColorMap[color]

	SpacesOfColor[inputLineIdx] += len(colorCode) + len(Reset)

	return fmt.Sprintf("%s%s%s", colorCode, s, Reset)
}

// Checks if the given index is within any of the ranges specified by the start indices and length
func isInRange(index int, indices []int, length int) bool {
	for _, start := range indices {
		if index >= start && index < start+length {
			return true
		}
	}
	return false
}

// FindSubStringIndices returns the start indices of all occurrences of the substring in the main string
func FindSubStringIndices(str, substr string) []int {
	var indices []int
	for i := 0; i < len(str); i++ {
		if strings.HasPrefix(str[i:], substr) {
			indices = append(indices, i)
			if len(str) > i+len(substr) {
				i = i + len(substr) - 1
			}
		}
	}
	return indices
}
