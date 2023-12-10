package day06

import "testing"

func TestPart1Example1(t *testing.T) {
	result, err := Part1("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 288 {
		t.Errorf("Expected 288, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if result != 1084752 {
		t.Errorf("Expected 1084752, got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	result, err := Part2("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 71503 {
		t.Errorf("Expected 71503, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 28228952 {
		t.Errorf("Expected 28228952, got %d", result)
	}
}
