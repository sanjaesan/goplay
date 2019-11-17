package goplay

// Stack -
type Stack struct {
	queueOne *queue
	queueTwo *queue
}

// NewStack -
func newStack() *Stack {
	return &Stack{
		queueOne: newQueue(),
		queueTwo: newQueue(),
	}
}

func (s *Stack) push(x int) {
	if s.queueOne.Len() == 0{
		s.queueOne, s.queueTwo = s.queueTwo, s.queueOne
	}
	s.queueOne.push(x)
}

func (s *Stack) pop() int {
	// --- ---
	return 0
}

type queue struct {
	numbers []int
}

func newQueue() *queue {
	return &queue{
		numbers: []int{},
	}
}

func (s *queue) push(x int) {
	s.numbers = append(s.numbers, x)
}

func (s *queue) pop() int {
	ret := s.numbers[0]
	s.numbers = s.numbers[1:]
	return ret
}

func (s *queue) Len() int {
	return len(s.numbers)
}
func (s *queue) IsEmpty() bool {
	return s.Len() == 0
}