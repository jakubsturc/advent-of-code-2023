package day01

import (
	"testing"
)

func TestPart1Example1(t *testing.T) {
	result, err := Part1("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 142 {
		t.Errorf("Expected 142, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 54081 {
		t.Errorf("Expected 54081, got %d", result)
	}
}

func TestPart2Example2(t *testing.T) {
	result, err := Part2("./example2.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 281 {
		t.Errorf("Expected 281, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 54649 {
		t.Errorf("Expected 54649, got %d", result)
	}
}
