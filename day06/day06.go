package day06

import (
	"aoc2023/shared"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Part1(path string) (uint64, error) {
	times, distances, err := parse1(path)
	if err != nil {
		return 0, err
	}

	result := find(times, distances)

	var product = uint64(1)
	for _, r := range result {
		product *= r
	}
	return product, nil
}

func Part2(path string) (uint64, error) {
	time, distance, _ := parse2(path)
	results := find([]uint64{time}, []uint64{distance})
	return results[0], nil
}

func find(times []uint64, distances []uint64) []uint64 {
	l := uint64(len(times))
	result := make([]uint64, l)

	for i := uint64(0); i < l; i++ {
		time := times[i]
		distance := distances[i]
		for charge := uint64(0); charge < time; charge++ {
			succeed := eval(time, charge, distance)
			if succeed {
				result[i]++
			}
		}
	}
	return result
}

func eval(time uint64, charge uint64, distance uint64) bool {
	possible := (time - charge) * charge
	succeed := possible > distance
	return succeed
}

func parse1(path string) ([]uint64, []uint64, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return nil, nil, err
	}
	times, err := parseNumbers(lines[0])
	if err != nil {
		return nil, nil, err
	}
	distances, err := parseNumbers(lines[1])
	if err != nil {
		return nil, nil, err
	}

	if len(times) != len(distances) {
		return nil, nil, errors.New("times and distances must be the same length")
	}

	return times, distances, nil
}

func parse2(path string) (uint64, uint64, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, 0, err
	}
	time, err := parseNumber(lines[0])
	if err != nil {
		return 0, 0, err
	}
	distance, err := parseNumber(lines[1])
	if err != nil {
		return 0, 0, err
	}
	return time, distance, nil
}

func parseNumbers(s string) ([]uint64, error) {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	numbers := make([]uint64, len(matches))
	for i, match := range matches {
		num, err := strconv.Atoi(match)
		if err != nil {
			return nil, err
		}
		numbers[i] = uint64(num)
	}
	return numbers, nil
}

func parseNumber(s string) (uint64, error) {
	idx := strings.Index(s, ":")
	if idx == -1 {
		return 0, errors.New("invalid input")
	}
	s = s[idx+1:]
	s = strings.ReplaceAll(s, " ", "")
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
