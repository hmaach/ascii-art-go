package asciiart

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Alignments = []string{
	"left", "center", "right", "justify",
}

func IsValidAlignment(align string) bool {
	for _, a := range Alignments {
		if a == align {
			return true
		}
	}
	return false
}

func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error getting terminal width: %v\n", err)
		os.Exit(1)
	}

	sliceOutput := strings.Fields(string(output))
	width, errAtoi := strconv.Atoi(sliceOutput[1])
	if errAtoi != nil {
		fmt.Printf("Error: %v\n", errAtoi)
		os.Exit(1)
	}

	return width
}

func AddSpacesBeforeLines(
	lines []string,
	align string,
	width, outpuLength int,
) string {
	var spacesToAdd int
	switch align {
	case "center":
		spacesToAdd = (width - outpuLength + (SpacesOfColor / 8)) / 2
	case "right":
		spacesToAdd = width - outpuLength + (SpacesOfColor / 8)
	}

	if spacesToAdd < 0 {
		spacesToAdd = 0
	}
	spaceString := strings.Repeat(" ", spacesToAdd)

	for i, line := range lines {
		if i < len(lines)-1 {
			lines[i] = spaceString + line
		}
	}
	return strings.Join(lines, "\n")
}

func AddSpacesBetweenWords(
	lines []string,
	width, outpuLength int,
) string {
	wordsNumber := strings.Count(lines[0], "{space}") + 1
	if wordsNumber > 1 {
		wordsWithoutSpaces := strings.ReplaceAll(lines[0], "{space}", "")
		lettersLength := len(wordsWithoutSpaces) + 1
		spacesToAdd := (width - lettersLength + (SpacesOfColor / 8)) / (wordsNumber - 1)
		if spacesToAdd < 0 {
			spacesToAdd = 0
		}

		spaceString := strings.Repeat(" ", spacesToAdd)

		for i, line := range lines {
			if i < len(lines)-1 {
				lines[i] = strings.ReplaceAll(line, "{space}", spaceString)
			}
		}
	}
	return strings.Join(lines, "\n")
}

func Justify(str, align string) string {
	var result string
	// check if the alignment is valid
	if !IsValidAlignment(align) {
		PrintAlignments()
	}

	width := GetTerminalWidth()

	lines := strings.Split(str, "\n")
	outpuLength := len(lines[0])
	if align == "justify" {
		result = AddSpacesBetweenWords(lines, width, outpuLength)
	} else {
		result = AddSpacesBeforeLines(lines, align, width, outpuLength)
	}
	return result
}
