package ds

// Stack is a generic stack data structure that can hold any type T.
// It provides methods to push, pop, and peek elements from the stack.
type Stack[T any] struct {
	items []T
}

// Push adds the provided item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack. If the stack is
// empty, it panics with the message "The stack is empty".
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("The stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Peek returns the top element of the stack without removing it.
// If the stack is empty, it panics with the message "The stack is empty".
func (s *Stack[T]) Peek() T {
	if len(s.items) == 0 {
		panic("The stack is empty")
	}
	return s.items[len(s.items)-1]
}

// NewStack creates a new empty stack that can hold elements of type T.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}
