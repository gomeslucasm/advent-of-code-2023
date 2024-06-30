package utils

import (
	"bufio"
	"os"
	"path/filepath"
)

func GetTextLines(filePath string) ([]string, error) {
	relativePath := filePath

	absolutePath, err := filepath.Abs(relativePath)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(absolutePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
