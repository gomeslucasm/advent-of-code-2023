package main

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
)

func findNumberInLine(line string) int {

	re := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	numbers := re.FindAllString(line, -1)

	if len(numbers) == 0 {

		return 0
	}

	numbersMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	stringNumber := numbersMap[numbers[0]] + numbersMap[numbers[len(numbers)-1]]

	fmt.Println("numbers = ", numbers, " | string number = ", stringNumber, " | line = ", line)

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
	lines, err := utils.GetTextLines("challenge-1/2/data.txt")

	if err != nil {
		return 0
	}

	var calibration int = 0

	for _, line := range lines {
		calibration = calibration + findNumberInLine(line)
	}

	/* fmt.Println("calibValues = ", calibValues) */

	return calibration
}
