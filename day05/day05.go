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

	seedStarts, filters, err := parse(lines)
	if err != nil {
		return 0, err
	}

	seeds := make([]interval, len(seedStarts))
	for i, seedStart := range seedStarts {
		seeds[i] = interval{seedStart, 1}
	}

	return solve(seeds, filters)
}

func Part2(path string) (uint64, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	seedStarts, filters, err := parse(lines)
	if err != nil {
		return 0, err
	}

	seeds := make([]interval, len(seedStarts)/2)
	for i := 0; i < len(seedStarts); i += 2 {
		seeds[i/2] = interval{seedStarts[i], seedStarts[i+1]}
	}

	return solve(seeds, filters)
}

func solve(seeds []interval, filters []filter) (uint64, error) {
	min, hasMin := uint64(0), false

	for _, seed := range seeds {
		current := []interval{seed}
		for _, filter := range filters {
			current = filter.apply(current)
		}
		for _, n := range current {
			if !hasMin || n.start < min {
				min = n.start
				hasMin = true
			}
		}
	}

	if !hasMin {
		return 0, errors.New("no min")
	}

	return min, nil
}

func (f filter) apply(ns []interval) []interval {
	var res []interval
	for _, n := range ns {
		var mapped []interval // numbers which has been successfully mapped
		for _, scope := range f.scopes {
			source := interval{scope.source, scope.length}
			found, val := source.intersection(n)
			if !found {
				continue
			}
			mapped = append(mapped, val)
			start := scope.destination + val.start - scope.source
			destination := interval{start, val.length}
			res = append(res, destination)
		}
		notMapped := n.subtractMany(mapped)
		for _, n := range notMapped {
			res = append(res, n)
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
