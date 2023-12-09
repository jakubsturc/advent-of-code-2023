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

func TestPart1Example2(t *testing.T) {
	result, err := Part1("./input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if result != 218513636 {
		t.Errorf("Expected 218513636, got %d", result)
	}
}
