package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name     string
		push     []int
		expected []int
	}{
		{"Popping from an empty stack", []int{}, []int{}},
		{"Pushing and popping elements", []int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{}
			for _, v := range tt.push {
				s.Push(v)
			}

			for i, v := range tt.expected {
				result := s.Pop()
				if result != v {
					t.Errorf("at pop %d expected %d, got %d", i, v, result)
				}
			}
		})
	}
}
