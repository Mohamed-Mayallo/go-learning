package main

func main() {

}

type Stack[T int | string] struct {
	values []T
}

func (s *Stack[T]) Push(v T) {
	s.values = append(s.values, v)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	lastIndex := len(s.values) - 1
	el := s.values[lastIndex]
	s.values = s.values[:lastIndex]

	return el, true
}
