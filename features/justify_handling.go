package asciiart

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Alignments = []string{
	"center", "left", "right", "justify",
}

func IsValidAlignment(align string) bool {
	for _, a := range Alignments {
		if a == align {
			return true
		}
	}
	return false
}

func GetTerminalWidth() (int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	sliceOutput := strings.Fields(string(output))
	width, errAtoi := strconv.Atoi(sliceOutput[1])
	if errAtoi != nil {
		return 0, err
	}
	return width, nil
}
