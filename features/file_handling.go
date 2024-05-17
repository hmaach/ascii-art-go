package asciiart

import (
	"fmt"
	"os"
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
