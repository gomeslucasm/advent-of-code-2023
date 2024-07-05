package main

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"
)

func main() {
	fmt.Println("SOLUTION = ", solve())
}

func getPointLine(line string) int {
	data := strings.Split(line, ":")

	data = strings.Split(data[1], "|")

	re := regexp.MustCompile(`\d+`)

	solution := re.FindAllString(data[0], -1)

	numbers := re.FindAllString(data[1], -1)

	count := 0

	for _, number := range numbers {
		if slices.Contains(solution, number) {
			count += 1
		}
	}

	fmt.Println(solution, "|", numbers, "|", count)

	return int(math.Pow(2, float64(count-1)))
}

func solve() int {
	lines, err := utils.GetTextLines("challenge-4/1/data.txt")

	if err != nil {
		return 0
	}

	count := 0

	for _, line := range lines {
		count += getPointLine(line)
	}

	return count
}
