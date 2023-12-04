package day04

type set []int

func createFrom(numbers []int) set {
	return set(numbers)
}

func (s set) len() int {
	return len(s)
}

func (s set) contains(n int) bool {
	for _, i := range s {
		if i == n {
			return true
		}
	}
	return false
}

func (s set) intersect(other set) set {
	var result set
	for _, n := range s {
		if other.contains(n) {
			result = append(result, n)
		}
	}
	return result
}
