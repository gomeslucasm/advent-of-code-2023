package main

import (
	"fmt"
	"regexp"
	"strconv"
	"adventOfCode/utils"
)

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
	lines, err := utils.GetTextLines("1/data.txt")

	if err != nil {
		return 0
	}

	var calibration int = 0

	for i := 0; i < len(lines); i++ {

		calibration = calibration + findNumberInLine(lines[i])
	}

	return calibration
}
