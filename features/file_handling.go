package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func ReadBanner(banner string) map[rune][]string {
	data, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	stringData := string(data[1:])
	if banner == "thinkertoy" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}
	content := strings.Split(stringData, "\n\n")
	characterMatrix := ConvertToCharacterMatrix(content)
	return characterMatrix
}

func SaveFile(fileName string, str string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %v", fileName, err)
	}
	defer file.Close()

	data := []byte(str)
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file '%s': %v", fileName, err)
	}
	return nil
}
