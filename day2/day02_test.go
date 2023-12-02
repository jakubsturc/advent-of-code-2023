package day02

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	result, err := Part1("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 2278 {
		t.Errorf("Expected 2278, got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	result, err := Part2("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 2286 {
		t.Errorf("Expected 2286, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 67953 {
		t.Errorf("Expected 67953, got %d", result)
	}
}
