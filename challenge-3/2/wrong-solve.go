package main

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("SOLUTION = ", solve())
}

type NumberPosition struct {
	Value int
	Start int
	End   int
}

func findAsteriskPosition(line string) int {
	return strings.Index(line, "*")
}

func findNumberPositions(s string) []NumberPosition {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(s, -1)

	var positions []NumberPosition
	for _, match := range matches {
		value := s[match[0]:match[1]]
		intValue, _ := strconv.Atoi(value)
		positions = append(positions, NumberPosition{
			Value: intValue,
			Start: match[0],
			End:   match[1] - 1,
		})
	}

	return positions
}

func findNumbersNearestAsterisk(lines []string, asteriskPosition int) []NumberPosition {

	var numbersNearToAsterisk []NumberPosition

	for _, line := range lines {
		numbersPostitions := findNumberPositions(line)

		for _, numberPosition := range numbersPostitions {

			if (numberPosition.Start >= asteriskPosition-1 && numberPosition.Start <= asteriskPosition+1) || (numberPosition.End >= asteriskPosition-1 && numberPosition.End <= asteriskPosition+1) {
				numbersNearToAsterisk = append(numbersNearToAsterisk, numberPosition)
			}

		}
	}

	return numbersNearToAsterisk
}

func solve() int {
	lines, err := utils.GetTextLines("challenge-3/2/data.txt")

	if err != nil {
		return 0
	}

	count := 0

	for idx, line := range lines {

		asteriskPosition := findAsteriskPosition(line)

		if asteriskPosition < 0 {
			continue
		}

		var linesToTest []string

		if idx == 0 {
			linesToTest = append(linesToTest, line, lines[idx+1])
		} else if idx == len(lines)-1 {
			linesToTest = append(linesToTest, line, lines[idx-1])
		} else {
			linesToTest = append(linesToTest, line, lines[idx-1])
		}

		numbersNearToAsterisk := findNumbersNearestAsterisk(linesToTest, asteriskPosition)

		if len(numbersNearToAsterisk) == 2 {
			count += numbersNearToAsterisk[0].Value * numbersNearToAsterisk[1].Value
		}

		/* if len(numbersNearToAsterisk) == 2 {

			mult := 1
			for _, num := range numbersNearToAsterisk {
				mult = mult * num.Value
			}

			count += mult
		} */

	}

	return count
}
