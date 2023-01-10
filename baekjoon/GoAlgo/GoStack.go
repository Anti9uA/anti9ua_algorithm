package goStack

type Stack []int

// Push adds an element to the top of the stack
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack
func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
