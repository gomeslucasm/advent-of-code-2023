package main

import (
	"adventOfCode/matrix"
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("SOLUTION = ", solve())
}

func registerNumbersInLine(line string, lineIdx int) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllStringIndex(line, -1)

	for _, match := range matches {
		value := line[match[0]:match[1]]
		intValue, _ := strconv.Atoi(value)

		element := matrix.CreateMatrixElement(intValue, "number")

		for idx := match[0]; idx <= match[1]-1; idx++ {
			point := matrix.CreateMatrixPoint(*element, string(line[idx]), idx, lineIdx)
			matrix.RegisterMatrixPoint(*point)
		}
	}
}

func registerAsteriskInLine(line string, lineIdx int) {

	re := regexp.MustCompile(`\*`)
	matches := re.FindAllStringIndex(line, -1)

	for _, match := range matches {
		el := matrix.CreateMatrixElement("*", "asterisk")
		point := matrix.CreateMatrixPoint(*el, "*", match[0], lineIdx)

		matrix.RegisterMatrixPoint(*point)
	}
}

func solve() int {
	lines, err := utils.GetTextLines("challenge-3/2/data.txt")

	if err != nil {
		return 0
	}

	for idx, line := range lines {
		registerNumbersInLine(line, idx)
		registerAsteriskInLine(line, idx)
	}

	asteriskMatrixElements := matrix.ListMatrixPointsByType("asterisk")

	count := 0

	for _, v := range asteriskMatrixElements {

		uniqNeighbors := make([]matrix.MatrixElement, 0)

		for _, n := range v.FindNeighbors() {

			/* To not save repeated Neighbors */
			if !slices.Contains(uniqNeighbors, n.Element) {
				uniqNeighbors = append(uniqNeighbors, n.Element)
			}
		}

		mult := 1

		if len(uniqNeighbors) > 1 {
			for _, el := range uniqNeighbors {
				switch v := el.Value.(type) {
				case int:
					mult = mult * v
				}
			}
			count += mult
		}

	}

	return count
}
