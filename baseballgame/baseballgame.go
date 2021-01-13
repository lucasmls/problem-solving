package baseballgame

import (
	"errors"
	"strconv"
)

// Stack ...
type Stack struct {
	items []int
}

// Push ...
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

// Pop ...
func (s *Stack) Pop() (int, error) {
	if len(s.items) > 0 {
		lastItem := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]

		return lastItem, nil
	}

	return -1, errors.New("The stack is empty")
}

// Peek ...
func (s *Stack) Peek() (int, error) {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1], nil
	}

	return -1, errors.New("The stack is empty")
}

// Size ...
func (s *Stack) Size() int {
	return len(s.items)
}

func calPoints(ops []string) int {
	stack := Stack{}

	for i := 0; i < len(ops); i++ {
		op := string(ops[i])

		if op == "+" {
			sum := stack.items[stack.Size()-1] + stack.items[stack.Size()-2]
			stack.Push(int(sum))
			continue
		}

		if op == "D" {
			top, _ := stack.Peek()

			double := top * 2
			stack.Push(int(double))
			continue
		}

		if op == "C" {
			stack.Pop()
			continue
		}

		score, _ := strconv.Atoi(op)
		stack.Push(score)
	}

	result := 0
	for stack.Size() > 0 {
		top, _ := stack.Pop()
		result += top
	}

	return result
}
