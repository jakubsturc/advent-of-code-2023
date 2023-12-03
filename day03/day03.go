package day03

import (
	"aoc2023/shared"
	"strconv"
)

func Part1(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}
	parts := getPartNumbers(lines)
	sum := sum(parts)
	return sum, nil
}

func Part2(path string) (int, error) {
	return 0, nil
}

func getPartNumbers(lines []string) []int {
	maxX := len(lines[0])
	maxY := len(lines)
	parts := make([]int, 0)
	var current string = ""
	var isNextToSymbol bool = false
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			c := lines[y][x]
			if '0' <= c && c <= '9' {
				current += string(c)
				isNextToSymbol = isNextToSymbol || hasSymbolNeighbour(lines, x, y)
				continue
			}

			if current != "" {
				if isNextToSymbol {
					part, err := strconv.Atoi(current)
					if err != nil {
						panic(err)
					}
					parts = append(parts, part)
				}
				current = ""
				isNextToSymbol = false
			}
		}
		if current != "" {
			if isNextToSymbol {
				part, err := strconv.Atoi(current)
				if err != nil {
					panic(err)
				}
				parts = append(parts, part)
			}
			current = ""
			isNextToSymbol = false
		}
	}
	return parts
}

func hasSymbolNeighbour(lines []string, x, y int) bool {
	neighbours := getNeighbours(lines, x, y)
	for _, neighbour := range neighbours {
		if isSymbol(neighbour) {
			return true
		}
	}
	return false
}

func getNeighbours(lines []string, x, y int) []rune {
	var neighbours []rune
	maxX := len(lines[0])
	maxY := len(lines)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			xx := x + j
			yy := y + i
			if xx < 0 || xx >= maxX || yy < 0 || yy >= maxY {
				continue
			}
			neighbours = append(neighbours, rune(lines[yy][xx]))
		}
	}
	return neighbours
}

func isNumber(c rune) bool {
	return '0' <= c && c <= '9'
}

func isSymbol(c rune) bool {
	if c == '.' {
		return false
	}
	if isNumber(c) {
		return false
	}
	return true
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
