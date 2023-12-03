package day03

import (
	"aoc2023/shared"
	"strconv"
)

type symbol struct {
	x, y int
	c    rune
}

type number struct {
	sx, ex, yy int
	val        int
}

type gear struct {
	x, y int
	nums []number
}

type plan struct{ lines []string }

func Part1(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	// get numbers and symbols
	var p = plan{lines}
	numbers := p.getNumbers()
	symbols := p.getSymbols()

	// filter parts
	var parts []number
	for _, number := range numbers {
		for _, symbol := range symbols {
			if isNextTo(number, symbol) {
				parts = append(parts, number)
			}
		}
	}

	// sum part numbers
	var sum = 0
	for _, part := range parts {
		sum += part.val
	}

	return sum, nil
}

func Part2(path string) (int, error) {
	lines, err := shared.ReadAllLines(path)
	if err != nil {
		return 0, err
	}

	// get numbers and symbols
	var p = plan{lines}
	numbers := p.getNumbers()
	symbols := p.getSymbols()

	// filter gears
	var gears []gear
	for _, symbol := range symbols {
		if isStar(symbol.c) {
			gears = append(gears, gear{symbol.x, symbol.y, nil})
		}
	}

	// associate gears with numbers
	for i, gear := range gears {
		for _, number := range numbers {
			if isNextTo(number, symbol{gear.x, gear.y, '*'}) {
				gears[i].add(number)
			}
		}
	}

	// sum filter correct gears and sum their ratios
	var sum = 0
	for _, gear := range gears {
		if len(gear.nums) == 2 {
			sum += gear.nums[0].val * gear.nums[1].val
		}
	}

	return sum, nil
}

func isNextTo(n number, s symbol) bool {
	return n.yy-1 <= s.y &&
		s.y <= n.yy+1 &&
		n.sx-1 <= s.x &&
		s.x <= n.ex+1
}

func (g *gear) add(n number) {
	g.nums = append(g.nums, n)
}

func (p *plan) getSymbols() []symbol {
	var symbols []symbol
	for y, line := range p.lines {
		for x, c := range line {
			if isSymbol(c) {
				symbols = append(symbols, symbol{x, y, c})
			}
		}
	}
	return symbols
}

func (p *plan) getNumbers() []number {
	var numbers []number
	var current = ""
	var sx, ex, yy = 0, 0, 0

	yield := func() {
		if current != "" {
			num, err := strconv.Atoi(current)
			if err != nil {
				panic(err)
			}
			number := number{sx, ex, yy, num}
			numbers = append(numbers, number)
			current = ""
		}
	}

	for y, line := range p.lines {
		for x, c := range line {
			if isDigit(c) {
				if current == "" {
					sx = x
					yy = y
				}
				ex = x
				current += string(c)
				continue
			}
			yield()
		}
		yield()
	}
	return numbers
}

func isStar(c rune) bool {
	return c == '*'
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isDot(c rune) bool {
	return c == '.'
}

func isSymbol(c rune) bool {
	if isDot(c) {
		return false
	}
	if isDigit(c) {
		return false
	}
	return true
}
