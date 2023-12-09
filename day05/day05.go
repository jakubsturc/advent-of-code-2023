package day05

import (
	"aoc2023/shared"
	"errors"
	"strconv"
	"strings"
)

type scope struct{ destination, source, length uint64 }
type filter struct{ scopes []scope }

func Part1(path string) (uint64, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	seeds, filters, err := parse(lines)

	if err != nil {
		return 0, err
	}

	var locations []uint64
	for _, seed := range seeds {
		current := seed
		for _, filter := range filters {
			current = filter.apply(current)
		}
		locations = append(locations, current)
	}

	result := min64(locations)

	return result, nil
}

func (f filter) apply(value uint64) uint64 {
	for _, scope := range f.scopes {
		if value >= scope.source && value < scope.source+scope.length {
			return value + scope.destination - scope.source
		}
	}
	return value
}

func min64(numbers []uint64) uint64 {
	var res uint64
	for i, n := range numbers {
		if i == 0 || n < res {
			res = n
		}
	}
	return res
}

func parse(lines []string) ([]uint64, []filter, error) {
	if len(lines) == 0 {
		return nil, nil, errors.New("no lines")
	}

	seeds, err := parseSeeds(lines[0])
	if err != nil {
		return nil, nil, err
	}

	var filters []filter
	var scopes []scope

	for _, line := range lines[2:] {
		if line == "" {
			filters = append(filters, filter{scopes})
			scopes = nil
			continue
		}
		if strings.HasSuffix(line, ":") {
			continue
		}
		scope, err := parseScope(line)
		if err != nil {
			return seeds, filters, err
		}
		scopes = append(scopes, scope)
	}

	filters = append(filters, filter{scopes})
	return seeds, filters, nil
}

func parseScope(line string) (scope, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		return scope{}, errors.New("invalid line")
	}
	destination, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return scope{}, err
	}
	source, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return scope{}, err
	}
	length, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return scope{}, err
	}
	return scope{destination, source, length}, nil
}

func parseSeeds(str string) ([]uint64, error) {
	parts := strings.Split(str, ": ")
	if parts[0] != "seeds" {
		return nil, errors.New("invalid seeds")
	}
	parts = strings.Split(parts[1], " ")
	var seeds []uint64
	for _, part := range parts {
		seed, err := strconv.ParseUint(part, 10, 64)
		if err != nil {
			return nil, err
		}
		seeds = append(seeds, seed)
	}
	return seeds, nil
}
