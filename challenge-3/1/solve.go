package main

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type NumberPosition struct {
	Value int
	Start int
	End   int
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

func checkHasSymbol(str string) bool {
	symbols := "!@#$%^&*()-_=+[]{}|;:'\",<>/?"
	return strings.Contains(symbols, str)
}

func checkSymbolInLine(line string, numberPosition NumberPosition) bool {
	startPos := numberPosition.Start
	if startPos != 0 {
		startPos = numberPosition.Start - 1
	}

	endPos := numberPosition.End
	if endPos < len(line)-1 {
		endPos = numberPosition.End + 2
	}

	slice := line[startPos:endPos]

	for _, char := range slice {
		if checkHasSymbol(string(char)) {
			return true
		}
	}

	return false
}

func checkNearestSymbol(line string, adjacentLines []string, numberPosition NumberPosition) bool {

	if numberPosition.Start > 0 && checkHasSymbol(string(line[numberPosition.Start-1])) {
		return true
	}

	if numberPosition.End < len(line)-1 && checkHasSymbol(string(line[numberPosition.End+1])) {
		return true
	}

	for _, adjacentLine := range adjacentLines {
		if checkSymbolInLine(adjacentLine, numberPosition) {
			return true
		}
	}

	return false
}

func solve() int {
	lines, err := utils.GetTextLines("challenge-3/1/data.txt")

	if err != nil {
		return 0
	}

	count := 0

	for idx, line := range lines {
		numberPositions := findNumberPositions(line)

		for _, numberPosition := range numberPositions {

			hasSymbol := false

			if idx == 0 {
				hasSymbol = checkNearestSymbol(line, []string{lines[idx+1]}, numberPosition)
			} else if idx == len(line)-1 {
				hasSymbol = checkNearestSymbol(line, []string{lines[idx-1]}, numberPosition)
			} else {
				hasSymbol = checkNearestSymbol(line, []string{lines[idx-1], lines[idx+1]}, numberPosition)
			}

			if hasSymbol {
				count += numberPosition.Value
			}
		}

	}

	return count
}

func main() {
	fmt.Println("SOLUTION = ", solve())
}
