package asciiart

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadBanner(banner string) string {
	data, err := os.ReadFile("banners/" + banner)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
	stringData := string(data[1:])
	return stringData
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

func ValidateFileExtension(filename string) error {
	acceptableExtensions := []string{".txt", ".json"}
	extension := strings.ToLower(filepath.Ext(filename))
	if extension == "" {
		return fmt.Errorf("please use one of the following extensions for the output file: .txt")
	}
	for _, ext := range acceptableExtensions {
		if extension == ext {
			return nil
		}
	}
	return fmt.Errorf("invalid file extension '%s' for --output option. Please use one of the following extensions: .txt", extension)
}

