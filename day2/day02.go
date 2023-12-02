package day02

import (
	"aoc2023/shared"
	"errors"
	"strconv"
	"strings"
)

func Part1(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}
	var sum int = 0
	for _, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		if isValid(game, 12, 13, 14) {
			sum += game.id
		}
	}
	return sum, nil
}

func Part2(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	sum := 0

	for _, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			return 0, err
		}
		sum += power(game)
	}
	return sum, nil
}

type draw struct {
	r int
	g int
	b int
}

type game struct {
	id    int
	draws []draw
}

func isValid(game game, r, g, b int) bool {
	for _, draw := range game.draws {
		if draw.r > r || draw.g > g || draw.b > b {
			return false
		}
	}
	return true
}

func power(game game) int {
	r, g, b := 0, 0, 0
	for _, draw := range game.draws {
		if draw.r > r {
			r = draw.r
		}
		if draw.g > g {
			g = draw.g
		}
		if draw.b > b {
			b = draw.b
		}
	}
	return r * g * b
}

func parseGame(line string) (game, error) {
	parts := strings.Split(line, ": ")
	id, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		return game{}, err
	}

	draws, err := parseDraws(parts[1])
	if err != nil {
		return game{}, err
	}
	return game{id, draws}, nil
}

func parseDraws(str string) ([]draw, error) {
	parts := strings.Split(str, "; ")
	draws := make([]draw, len(parts))
	for i, part := range parts {
		draw, err := parseDraw(part)
		if err != nil {
			return nil, err
		}
		draws[i] = draw
	}
	return draws, nil
}

func parseDraw(part string) (draw, error) {
	parts1 := strings.Split(part, ", ")
	var r, g, b int
	for _, part := range parts1 {
		parts2 := strings.Split(part, " ")
		v, err := strconv.Atoi(parts2[0])
		if err != nil {
			return draw{}, err
		}
		switch strings.TrimRight(parts2[1], "\r\n") {
		case "red":
			r = v
		case "green":
			g = v
		case "blue":
			b = v
		default:
			return draw{}, errors.New("Invalid color: " + parts2[1])
		}
	}
	return draw{r, g, b}, nil
}
