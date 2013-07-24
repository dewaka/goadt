package goadt_test

import (
	"testing"
	"github.com/dewaka/goadt"
)

func TestStackSimple(t *testing.T) {
	s := new(goadt.Stack)
	expect(t, s.Size(), 0)
	s.Push(3)
	expect(t, s.Size(), 1) 
	s.Pop()
	expect(t, s.Size(), 0) 
	s.Pop()
	expect(t, s.Size(), 0) 
}

func TestStackMultiple(t *testing.T) {
	const num = 100
	s := new(goadt.Stack)
	// Test a scenario where we insert 100 elements consecutively
	for i := 0; i < num; i++ {
		s.Push(i)
		// everytime we insert a new element new size should
		// be i+1
		expect(t, s.Size(), i+1) 
	}
	// Now test whether we can get back elements in LIFO order
	for i := num-1; i >= 0; i-- {
		n, m := s.Pop()
		expect(t, m, true) // Popping should succeed
		expect(t, n, i) // Popped value should be equal to i
	}

	// Now our stack should be empty
	expect(t, s.Empty(), true)
	expect(t, s.Size(), 0)
}
