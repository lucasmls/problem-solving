package stack

import "errors"

// Stack ...
type Stack struct {
	items []string
}

// Push ...
func (s *Stack) Push(value string) {
	s.items = append(s.items, value)
}

// Pop ...
func (s *Stack) Pop() (string, error) {
	if len(s.items) > 0 {
		lastItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]

		return lastItem, nil
	}

	return "", errors.New("The stack is empty")
}

// Peek ...
func (s *Stack) Peek() (string, error) {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1], nil
	}

	return "", errors.New("The stack is empty")
}

// ToString ...
func (s *Stack) ToString() string {
	str := ""

	for i := 0; i < len(s.items); i++ {
		char := s.items[i]
		str = str + char
	}

	return str
}
