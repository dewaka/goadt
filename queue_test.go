package goadt_test

import (
	"testing"
	"github.com/dewaka/goadt"
)

func TestQueueSimple(t *testing.T) {
	q := new(goadt.Queue)
	expect(t, q.Empty(), true) // a new queue should be empty
	expect(t, q.Size(), 0) // new queue should be size 0

	// Insert some elements
	q.Put(42)
	expect(t, q.Empty(), false) // now the queue should not be empty
	expect(t, q.Size(), 1) // new queue should be size 1
}

func TestQueueMultiple(t *testing.T) {
	const num = 100
	q := new(goadt.Queue)
	// Test a scenario where we insert 100 elements consecutively
	for i := 0; i < num; i++ {
		q.Put(i)
		// everytime we insert a new element new size should
		// be i+1
		expect(t, q.Size(), i+1) 
	}
	// Now test whether we can get back elements in FIFO order
	for i := 0; i < num; i++ {
		// check the size as we do operations
		expect(t, q.Size(), num - i)
		e := q.Remove()
		expect(t, e, i) // Removed value should be equal to i
	}

	// Now our stack should be empty
	expect(t, q.Empty(), true)
	expect(t, q.Size(), 0)
}
