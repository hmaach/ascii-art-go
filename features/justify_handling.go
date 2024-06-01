package asciiart

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

// Alignments contains the list of valid alignments
var Alignments = []string{
	"left", "center", "right", "justify",
}

// getTerminalWidth fetches the current terminal width
func getTerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error getting terminal width: %v", err)
	}

	sliceOutput := strings.Fields(string(output))
	if len(sliceOutput) < 2 {
		return 0, errors.New("unexpected output from stty command")
	}

	width, err := strconv.Atoi(sliceOutput[1])
	if err != nil {
		return 0, fmt.Errorf("error converting terminal width to integer: %v", err)
	}

	return width, nil
}

// add leading spaces to lines based on the specified spaces number
func addSpacesBeforeLines(
	lines []string,
	spacesToAdd int,
) string {
	spaceString := strings.Repeat(" ", spacesToAdd)

	for i, line := range lines {
		if i < len(lines)-1 {
			lines[i] = spaceString + line
		}
	}
	return strings.Join(lines, "\n")
}

// adds spaces between words to justify text alignment
func addSpacesBetweenWords(
	lines []string,
	width, outputWithoutSpaces int,
) string {
	spaceCountBetweenWords := strings.Count(lines[0], "{space}")
	spacesToAdd := (width - outputWithoutSpaces) / spaceCountBetweenWords
	spaceString := strings.Repeat(" ", spacesToAdd)

	for i := 0; i < len(lines); i++ {
		if i < len(lines)-1 {
			lines[i] = strings.ReplaceAll(lines[i], "{space}", spaceString)
		}
	}
	return strings.Join(lines, "\n")
}

// alignText aligns text based on the alignment type and terminal width.
func alignText(lines []string, flags map[string]string, width, LineIdx, outputLength int) string {
	align := flags["align"]
	spacesToAdd := (width - outputLength)

	if align == "center" {
		spacesToAdd /= 2
	}

	if align == "justify" {
		outputWithoutSpaces := len(strings.ReplaceAll(lines[0], "{space}", ""))
		// Adjust for color codes if present
		if flags["color"] != "" {
			outputWithoutSpaces -= SpacesOfColor[LineIdx] / 8 // Adjust length for color codes
		}

		return addSpacesBetweenWords(lines, width, outputWithoutSpaces)
	}

	return addSpacesBeforeLines(lines, spacesToAdd)
}

// Justify aligns the ASCII art string based on the specified alignment.
func Justify(str string, LineIdx int, flags map[string]string) (string, error) {
	// Validate alignment
	align := flags["align"]
	if !slices.Contains(Alignments, align) {
		return "", fmt.Errorf("invalid alignment: '%s'.\nValid alignments are: %v", align, Alignments)
	}

	width, err := getTerminalWidth()
	if err != nil {
		return "", err
	}

	lines := strings.Split(str, "\n")

	// Adjust output length for color codes if present
	outputLength := len(lines[0])
	if flags["color"] != "" {
		outputLength -= SpacesOfColor[LineIdx] / 8 // Adjust length for color codes
	}

	return alignText(lines, flags, width, LineIdx, outputLength), nil
}
