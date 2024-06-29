package main

import (
	"fmt"
	"strings"
	"adventOfCode/utils"
)

func main() {
	fmt.Println("SOLUTION = ", solve())
}

func parseLineInfo(line string) (int, map[string]int) {
	parts := strings.SplitN(line, ":", 2)

	cubesCount := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	var gameNumber int

	_, err := fmt.Sscanf(parts[0], "Game %d", &gameNumber)

	if err != nil {
		return 0, nil
	}

	for _, combination := range strings.Split(parts[1], ";") {
		for _, sentence := range strings.Split(combination, ",") {
			var number int
			var color string

			fmt.Sscanf(strings.TrimSpace(sentence), "%d %s", &number, &color)

			if number > cubesCount[color] {
				cubesCount[color] = number
			}
		}
	}

	return gameNumber, cubesCount
}

func solve() int {
	nRedCubes := 12
	nGreenCubes := 13
	nBlueCubes := 14

	count := 0

	lines, err := utils.GetTextLines("2/data.txt")

	if err != nil {
		return 0
	}

	for _, line := range lines {
		gameNumber, cubesCount := parseLineInfo(line)

		if cubesCount["red"] <= nRedCubes && cubesCount["green"] <= nGreenCubes && cubesCount["blue"] <= nBlueCubes {
			count = count + gameNumber
		}
	}

	return count
}
