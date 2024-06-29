package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func getTextLines() ([]string, error) {
	relativePath := "1/data.txt"

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
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func findNumberInLine(line string) int {

	re := regexp.MustCompile(`[0-9]`)
	numbers := re.FindAllString(line, -1)

	if len(numbers) == 0 {
		return 0
	}

	stringNumber := numbers[0] + numbers[len(numbers)-1]

	number, err := strconv.Atoi(stringNumber)

	if err != nil {
		return 0
	}

	return number

}

func main() {
	fmt.Println("Solution = ", solve())
}

func solve() int {
	lines, err := getTextLines()

	if err != nil {
		return 0
	}

	var calibration int = 0

	for i := 0; i < len(lines); i++ {

		calibration = calibration + findNumberInLine(lines[i])
	}

	return calibration
}
