package goadt_test

import (
	"fmt"
	"github.com/dewaka/goadt"
	"testing"
)

func TestAvlSimple(t *testing.T) {
	q := goadt.NewAVL(CInt(32))
	expect(t, q.Value(), CInt(32)) // new queue should be size 0
	expect(t, q.Contains(CInt(32)), true)

	q = q.Insert(CInt(7))
	expect(t, q.Contains(CInt(7)), true)

	expect(t, q.Contains(CInt(100)), false)
	q = q.Insert(CInt(100))
	expect(t, q.Contains(CInt(100)), true)
}

func TestAvlDepth(t *testing.T) {
	q := goadt.NewAVL(CInt(7))
	expect(t, q.Depth(), 1)

	q = q.Insert(CInt(10))
	expect(t, q.Depth(), 2)

	// less than 7 so to the left side without increasing the depth
	q = q.Insert(CInt(5))
	expect(t, q.Depth(), 2)

	// Just check all exists
	expect(t, q.Contains(CInt(7)), true)
	expect(t, q.Contains(CInt(10)), true)
	expect(t, q.Contains(CInt(5)), true)
}

func TestAVLRotation(t *testing.T) {
	q := goadt.NewAVL(CInt(3)).Insert(CInt(5)).Insert(CInt(2))
	expect(t, q.Contains(CInt(3)), true)
	expect(t, q.Contains(CInt(5)), true)
	expect(t, q.Contains(CInt(2)), true)

	p := q.RotateLeft()
	expect(t, p.Contains(CInt(3)), true)
	expect(t, p.Contains(CInt(5)), true)
	expect(t, p.Contains(CInt(2)), true)
}

func TestAVLToSlice(t *testing.T) {
	q := goadt.NewAVL(CInt(3)).Insert(CInt(5)).Insert(CInt(2))
	p := goadt.NewAVL(CInt(3)).Insert(CInt(5)).Insert(CInt(2))
	fmt.Println(q.AvlToSlice())

	q = q.RotateLeft()
	p = p.RotateRight()
	fmt.Println(q.AvlToSlice())
	fmt.Println(p.AvlToSlice())
}

func TestAvlRotation(t *testing.T) {
	q := goadt.NewAVL(CInt(34))
	q = q.Insert(CInt(35)).
		Insert(CInt(3)).
		Insert(CInt(43)).
		Insert(CInt(53)).
		Insert(CInt(23)).
		Insert(CInt(55))

	fmt.Println(q.AvlToSlice())
}
