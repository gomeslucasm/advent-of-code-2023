package main

import (
	"adventOfCode/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func findNumberInLine(line string) int {
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

	substringsMap := make(map[int]string)

	for key := range numbersMap {
		index := strings.Index(line, key)
		if index != -1 {
			substringsMap[index] = key
		}
	}

	if len(substringsMap) == 0 {
		return 0
	}

	keys := make([]int, 0, len(substringsMap))
	for k := range substringsMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	stringNumber := numbersMap[substringsMap[keys[0]]] + numbersMap[substringsMap[keys[len(keys)-1]]]

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

	calibration := 0

	for _, line := range lines {
		calibration += findNumberInLine(line)
	}

	return calibration
}
