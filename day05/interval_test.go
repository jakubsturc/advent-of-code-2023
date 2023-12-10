package day05

import "testing"

func TestIntersection_StrictSubset(t *testing.T) {
	a := interval{0, 10}
	b := interval{1, 8}

	result, intersection := a.intersection(b)

	if !result {
		t.Error("Expected intersection")
	}

	if intersection != b {
		t.Errorf("Expected %v, got %v", b, intersection)
	}
}

func TestIntersection_StrictSuperset(t *testing.T) {
	a := interval{1, 8}
	b := interval{0, 10}

	result, intersection := a.intersection(b)

	if !result {
		t.Error("Expected intersection")
	}

	if intersection != a {
		t.Errorf("Expected %v, got %v", a, intersection)
	}
}

func TestIntersection_NoIntersection(t *testing.T) {
	a := interval{0, 10}
	b := interval{11, 10}

	result, _ := a.intersection(b)

	if result {
		t.Error("Expected no intersection")
	}
}

func TestIntersection_PartialIntersection(t *testing.T) {
	a := interval{0, 10}
	b := interval{5, 10}

	result, intersection := a.intersection(b)

	if !result {
		t.Error("Expected intersection")
	}

	expected := interval{5, 5}
	if intersection != expected {
		t.Errorf("Expected %v, got %v", interval{5, 5}, intersection)
	}
}

func TestSubtract_NoIntersection(t *testing.T) {
	a := interval{0, 10}
	b := interval{11, 10}

	result := a.subtract(b)

	if len(result) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(result))
	}

	if result[0] != a {
		t.Errorf("Expected %v, got %v", a, result[0])
	}
}

func TestSubtract_PartialIntersectionEnd(t *testing.T) {
	a := interval{0, 10}
	b := interval{5, 10}

	result := a.subtract(b)

	if len(result) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(result))
	}

	expected := interval{0, 5}
	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestSubtract_PartialIntersectionStart(t *testing.T) {
	a := interval{5, 10}
	b := interval{0, 10}

	result := a.subtract(b)

	if len(result) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(result))
	}

	expected := interval{10, 5}
	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestSubtract_PartialIntersectionMiddle(t *testing.T) {
	a := interval{0, 10}
	b := interval{2, 2}

	result := a.subtract(b)

	if len(result) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(result))
	}

	expected := []interval{{0, 2}, {4, 6}}
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSubtract_FullIntersection(t *testing.T) {
	a := interval{0, 10}
	b := interval{0, 10}

	result := a.subtract(b)

	if len(result) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(result))
	}
}

func TestSubtractMany_NoIntersection(t *testing.T) {
	a := interval{0, 10}
	b := []interval{{11, 10}, {12, 10}}

	result := a.subtractMany(b)

	if len(result) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(result))
	}

	if result[0] != a {
		t.Errorf("Expected %v, got %v", a, result[0])
	}
}

func TestSubtractMany_PartialIntersectionEnd(t *testing.T) {
	a := interval{0, 10}
	b := []interval{{5, 10}, {6, 10}}

	result := a.subtractMany(b)

	if len(result) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(result))
	}

	expected := interval{0, 5}
	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestSubtractMany_PartialIntersectionStart(t *testing.T) {
	a := interval{5, 10}
	b := []interval{{0, 10}, {1, 10}}

	result := a.subtractMany(b)

	if len(result) != 1 {
		t.Errorf("Expected 1 interval, got %d", len(result))
	}

	expected := interval{11, 4}
	if result[0] != expected {
		t.Errorf("Expected %v, got %v", expected, result[0])
	}
}

func TestSubtractMany_PartialIntersectionMiddle(t *testing.T) {
	a := interval{0, 10}
	b := []interval{{2, 2}, {4, 2}}

	result := a.subtractMany(b)

	if len(result) != 2 {
		t.Errorf("Expected 2 intervals, got %d", len(result))
	}

	expected := []interval{{0, 2}, {6, 4}}
	if result[0] != expected[0] || result[1] != expected[1] {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSubtractMany_FullIntersection(t *testing.T) {
	a := interval{0, 10}
	b := []interval{{0, 2}, {2, 8}}

	result := a.subtractMany(b)

	if len(result) != 0 {
		t.Errorf("Expected 0 intervals, got %d", len(result))
	}
}
