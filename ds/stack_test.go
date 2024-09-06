package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOperations(t *testing.T) {
	tests := []struct {
		name     string
		pushVals []int
		popVals  []int
	}{
		{"Push and Pop multiple items", []int{1, 2, 3}, []int{3, 2, 1}},
		{"Push and Pop single item", []int{42}, []int{42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStack[int]()
			for _, v := range tt.pushVals {
				s.Push(v)
			}
			for _, expected := range tt.popVals {
				assert.Equal(t, expected, s.Pop())
			}
		})
	}

	t.Run("Pop from empty stack panics", func(t *testing.T) {
		s := NewStack[float64]()
		assert.Panics(t, func() { s.Pop() })
	})

	t.Run("Peek from empty stack panics", func(t *testing.T) {
		s := NewStack[bool]()
		assert.Panics(t, func() { s.Peek() })
	})
}
