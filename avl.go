package goadt

import "fmt" // for debugging time only

type AvlNode struct {
	val         Comparable
	left, right *AvlNode
}

func NewAVL(val Comparable) *AvlNode {
	return &AvlNode{val, nil, nil}
}

func (n *AvlNode) Value() Comparable {
	return n.val
}

func (n *AvlNode) Insert(val Comparable) *AvlNode {
	nn := insertAvl(n, &AvlNode{val, nil, nil})
	return nn
}

func (n *AvlNode) Contains(val Comparable) bool {
	switch n.val.Compare(val) {
	case 0:
		return true
	case -1: // node value is less than comparing value
		if n.right == nil {
			return false
		} else {
			return n.right.Contains(val)
		}
	case 1: //node value greater than comparing value
		if n.left == nil {
			return false
		} else {
			return n.left.Contains(val)
		}
	default:
		return false // we shouldn't be here anyway
	}
}

func insertAvl(n *AvlNode, m *AvlNode) *AvlNode {
	if n == nil {
		return m
	}
	if m == nil {
		return n
	}

	var nn *AvlNode

	switch m.val.Compare(n.val) {
	case 1: // right side
		nn = &AvlNode{n.val, n.left, insertAvl(n.right, m)}
	default: // left side
		nn = &AvlNode{n.val, insertAvl(n.left, m), n.right}
	}

	switch checkAvl(nn) {
	case 1:
		fmt.Println("Rotating to right")
		nn = rotateRight(nn)
	case -1:
		fmt.Println("Rotating to right")
		nn = rotateLeft(nn)
	default:
		fmt.Println("Don't have to rotate")
		// nothing to do
	}

	return nn
}

func (n *AvlNode) Depth() int {
	return depth(n)
}

func depth(n *AvlNode) int {
	if n == nil {
		return 0
	} else {
		return max(depth(n.left), depth(n.right)) + 1
	}
}

func (n *AvlNode) RotateLeft() *AvlNode {
	return rotateLeft(n)
}

func (n *AvlNode) RotateRight() *AvlNode {
	return rotateRight(n)
}

func rotateLeft(n *AvlNode) *AvlNode {
	q := n
	p := q.right
	c := q.left
	a := p.left
	b := p.right

	q = &AvlNode{q.val, c, a}
	p = &AvlNode{p.val, q, b}

	return p
}

func rotateRight(n *AvlNode) *AvlNode {
	q := n
	p := q.left
	c := q.right
	a := p.left
	b := p.right

	q = &AvlNode{q.val, b, c}
	p = &AvlNode{p.val, a, q}

	return p
}

func (n *AvlNode) AvlToSlice() []Comparable {
	res := make([]Comparable, 0)
	return appendAvlElemsToSlice(n, res)
}

func appendAvlElemsToSlice(node *AvlNode, res []Comparable) []Comparable {
	if node != nil {
		res = append(res, node.val)
		res = appendAvlElemsToSlice(node.left, res)
		res = appendAvlElemsToSlice(node.right, res)
		return res
	}
	return res
}

func checkAvl(n *AvlNode) int {
	d := depth(n.right) - depth(n.left)
	switch {
	case d >= 2:
		return 1
	case d <= -2:
		return -1
	default:
		return 0
	}
}
