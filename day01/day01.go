package day01

import (
	"aoc2023/shared"
	"strings"
)

func Part1(path string) (int, error) {
	return sum(path, false)
}

func Part2(path string) (int, error) {
	return sum(path, true)
}

func sum(path string, matchText bool) (int, error) {
	lines, err := shared.ReadAllLines(path)

	if err != nil {
		return 0, err
	}

	var sum int
	for _, line := range lines {
		var first = firstNumber(line, matchText)
		var last = lastNumber(line, matchText)
		sum += first*10 + last
	}

	return sum, nil
}

type numOcc struct {
	num int
	pos int
}

func firstNumber(line string, matchText bool) int {
	occs := getOccurrences(line, matchText, false)

	first := numOcc{num: -1, pos: -1}

	for _, occ := range occs {
		if occ.pos != -1 {
			if first.pos == -1 {
				first = occ
			} else if occ.pos < first.pos {
				first = occ
			}
		}
	}

	return first.num
}

func lastNumber(line string, matchText bool) int {
	occs := getOccurrences(line, matchText, true)

	last := numOcc{num: -1, pos: -1}

	for _, occ := range occs {
		if occ.pos != -1 {
			if last.pos == -1 {
				last = occ
			} else if occ.pos > last.pos {
				last = occ
			}
		}
	}

	return last.num
}

func getOccurrences(line string, matchText bool, last bool) []numOcc {
	occs := []numOcc{}
	occs = append(occs, numOcc{num: 1, pos: getOccurrence(line, "1", last)})
	occs = append(occs, numOcc{num: 2, pos: getOccurrence(line, "2", last)})
	occs = append(occs, numOcc{num: 3, pos: getOccurrence(line, "3", last)})
	occs = append(occs, numOcc{num: 4, pos: getOccurrence(line, "4", last)})
	occs = append(occs, numOcc{num: 5, pos: getOccurrence(line, "5", last)})
	occs = append(occs, numOcc{num: 6, pos: getOccurrence(line, "6", last)})
	occs = append(occs, numOcc{num: 7, pos: getOccurrence(line, "7", last)})
	occs = append(occs, numOcc{num: 8, pos: getOccurrence(line, "8", last)})
	occs = append(occs, numOcc{num: 9, pos: getOccurrence(line, "9", last)})
	occs = append(occs, numOcc{num: 0, pos: getOccurrence(line, "0", last)})

	if matchText {
		occs = append(occs, numOcc{num: 1, pos: getOccurrence(line, "one", last)})
		occs = append(occs, numOcc{num: 2, pos: getOccurrence(line, "two", last)})
		occs = append(occs, numOcc{num: 3, pos: getOccurrence(line, "three", last)})
		occs = append(occs, numOcc{num: 4, pos: getOccurrence(line, "four", last)})
		occs = append(occs, numOcc{num: 5, pos: getOccurrence(line, "five", last)})
		occs = append(occs, numOcc{num: 6, pos: getOccurrence(line, "six", last)})
		occs = append(occs, numOcc{num: 7, pos: getOccurrence(line, "seven", last)})
		occs = append(occs, numOcc{num: 8, pos: getOccurrence(line, "eight", last)})
		occs = append(occs, numOcc{num: 9, pos: getOccurrence(line, "nine", last)})
		occs = append(occs, numOcc{num: 0, pos: getOccurrence(line, "zero", last)})
	}

	return occs
}

func getOccurrence(line string, str string, last bool) int {
	if last {
		return strings.LastIndex(line, str)
	}
	return strings.Index(line, str)
}
