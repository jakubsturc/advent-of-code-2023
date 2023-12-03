package day03

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	result, err := Part1("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 4361 {
		t.Errorf("Expected 4361, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 526404 {
		t.Errorf("Expected 526404, got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	result, err := Part2("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 467835 {
		t.Errorf("Expected 467835, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 84399773 {
		t.Errorf("Expected 84399773, got %d", result)
	}
}
