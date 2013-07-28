package goadt_test

import (
	"github.com/dewaka/goadt"
	"testing"
	"fmt"
)

type CInt int

func (a CInt) Compare(b goadt.Comparable) int {
	c := b.(CInt)
	switch {
	case int(a) < int(c):
		return -1
	case int(a) > int(c):
		return 1
	default:
		return 0
	}
}

func TestBSTSimple(t *testing.T) {
	q := goadt.NewBST(CInt(7))
	expect(t, q.Count(), 1)

	q.Insert(CInt(3))
	expect(t, q.Count(), 2)
}

func TestBSTMultipleInsertions(t *testing.T) {
	q := goadt.NewBST(CInt(0))
	for i := 1; i < 10; i++ {
		q.Insert(CInt(i))
		expect(t, q.Count(), i+1)
	}

	// Test contains
	for i := 0; i < 10; i++ {
		expect(t, q.Contains(CInt(i)), true)
	}

	expect(t, q.Contains(CInt(10)), false)
}

func TestToSlice(t *testing.T) {
	q := goadt.NewBST(CInt(0))
	for i := 1; i < 10; i++ {
		q.Insert(CInt(i))
	}
	qs := q.ToSlice()
	es := []goadt.Comparable{CInt(0), CInt(1), CInt(2), CInt(3), CInt(4), CInt(5), CInt(6), CInt(7), CInt(8), CInt(9)}
	expect(t, fmt.Sprintf("%v", qs), fmt.Sprintf("%v", es)) // fast way to compare slices by converting them to strings

	// Now we are going to remove 3, 7, 4 from BST and check again
	q.Remove(CInt(3))
	q.Remove(CInt(7))
	q.Remove(CInt(4))

	// now the expected slice does not contain 3, 7 and also 4
	es = []goadt.Comparable{CInt(0), CInt(1), CInt(2), CInt(5), CInt(6), CInt(8), CInt(9)} 
	expect(t, fmt.Sprintf("%v", q.ToSlice()), fmt.Sprintf("%v", es))
}
