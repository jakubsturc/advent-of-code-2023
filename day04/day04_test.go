package day04

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	result, err := Part1("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 13 {
		t.Errorf("Expected 13, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 25571 {
		t.Errorf("Expected 25571, got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	result, err := Part2("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 30 {
		t.Errorf("Expected 30, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 8805731 {
		t.Errorf("Expected 8805731, got %d", result)
	}
}
