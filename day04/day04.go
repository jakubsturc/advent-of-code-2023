package day04

import (
	"aoc2023/shared"
	"strconv"
	"strings"
)

func Part1(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, line := range lines {
		winning, having, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		cnt := winning.intersect(having).len()
		if cnt > 0 {
			sum += 1 << uint(cnt-1)
		}
	}
	return sum, nil
}

func Part2(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	l := len(lines)
	cnts := initCounts(l)

	for i, line := range lines {
		winning, having, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		s := winning.intersect(having).len()
		for j := 1; i+j < l && j <= s; j++ {
			cnts[i+j] += cnts[i]
		}
	}

	sum := 0
	for _, cnt := range cnts {
		sum += cnt
	}
	return sum, nil
}

func initCounts(len int) []int {
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[i] = 1
	}
	return result
}

func parseLine(line string) (set, set, error) {
	// parse line into two sets Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	parts := strings.Split(line, ": ")
	parts = strings.Split(parts[1], "|")
	winning, err := parseSet(parts[0])
	if err != nil {
		return nil, nil, err
	}
	having, err := parseSet(parts[1])
	if err != nil {
		return nil, nil, err
	}
	return winning, having, nil
}

func parseSet(s string) (set, error) {
	parts := strings.Split(s, " ")
	var result set
	for _, part := range parts {
		if part == "" {
			continue
		}
		n, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}
