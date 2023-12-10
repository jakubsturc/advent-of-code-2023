package day05

import "testing"

func TestPart1Example1(t *testing.T) {
	result, err := Part1("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 35 {
		t.Errorf("Expected 35, got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if result != 218513636 {
		t.Errorf("Expected 218513636, got %d", result)
	}
}

func TestPart2Example1(t *testing.T) {
	result, err := Part2("./example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	if result != 46 {
		t.Errorf("Expected 46, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := Part2("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if result != 81956384 {
		t.Errorf("Expected 81956384, got %d", result)
	}
}
