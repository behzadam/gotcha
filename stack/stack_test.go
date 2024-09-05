package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackOperations(t *testing.T) {
	t.Run("Push and Pop multiple items", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		assert.Equal(t, 3, s.Pop())
		assert.Equal(t, 2, s.Pop())
		assert.Equal(t, 1, s.Pop())
	})

	t.Run("Peek without modifying stack", func(t *testing.T) {
		s := NewStack[string]()
		s.Push("hello")
		s.Push("world")

		assert.Equal(t, "world", s.Peek())
		assert.Equal(t, "world", s.Peek())
		assert.Equal(t, "world", s.Pop())
		assert.Equal(t, "hello", s.Pop())
	})

	t.Run("Pop from empty stack panics", func(t *testing.T) {
		s := NewStack[float64]()
		assert.Panics(t, func() { s.Pop() })
	})

	t.Run("Peek from empty stack panics", func(t *testing.T) {
		s := NewStack[bool]()
		assert.Panics(t, func() { s.Peek() })
	})

	t.Run("Push and Pop with different types", func(t *testing.T) {
		s1 := NewStack[int]()
		s1.Push(42)
		assert.Equal(t, 42, s1.Pop())

		s2 := NewStack[string]()
		s2.Push("test")
		assert.Equal(t, "test", s2.Pop())

		s3 := NewStack[[]int]()
		s3.Push([]int{1, 2, 3})
		assert.Equal(t, []int{1, 2, 3}, s3.Pop())
	})

	t.Run("Large number of push and pop operations", func(t *testing.T) {
		s := NewStack[int]()
		for i := 0; i < 1000; i++ {
			s.Push(i)
		}
		for i := 999; i >= 0; i-- {
			assert.Equal(t, i, s.Pop())
		}
	})
}
