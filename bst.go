package goadt

import (
	"fmt"
)

type Comparable interface {
	Compare(Comparable) int
}

type BSTNode struct {
	val         Comparable
	left, right *BSTNode
}

func NewBST(val Comparable) *BSTNode {
	return &BSTNode{val, nil, nil}
}

func (n *BSTNode) Insert(val Comparable) *BSTNode {
	insert(n, val)
	return n
}

func (n *BSTNode) Contains(val Comparable) bool {
	return contains(n, val)
}

func (n *BSTNode) Value() Comparable {
	return n.val
}

func (n *BSTNode) Count() int {
	lc, rc := 0, 0
	if n.left != nil {
		lc = n.left.Count()
	}
	if n.right != nil {
		rc = n.right.Count()
	}
	return 1 + lc + rc
}

func (n *BSTNode) Remove(val Comparable) (*BSTNode, bool) {
	return remove(n, val)
}

func (n *BSTNode) String() string {
	ls, rs := "()", "()"
	if n.left != nil {
		ls = n.left.String()
	}
	if n.right != nil {
		rs = n.right.String()
	}
	return fmt.Sprintf("%v (%s) (%s)", n.val, ls, rs)
}

func (n *BSTNode) Append(s *BSTNode) *BSTNode {
	appendToNode(n, s)
	return n
}

func (n *BSTNode) Height() int {
	return height(n)
}

func (n *BSTNode) ToSlice() []Comparable {
	res := make([]Comparable, 0)
	return appendElemsToSlice(n, res)
}

func (n *BSTNode) Min() Comparable {
	if n.left != nil {
		return n.left.Min()
	} else {
		return n.val
	}
}

func (n *BSTNode) Max() Comparable {
	if n.right != nil {
		return n.right.Max()
	} else {
		return n.val
	}
}

func insert(node *BSTNode, val Comparable) {
	switch {
	case node == nil:
		return
	case val.Compare(node.val) < 0:
		if node.left != nil {
			insert(node.left, val)
		} else {
			node.left = &BSTNode{val, nil, nil}
		}
	case val.Compare(node.val) >= 0:
		if node.right != nil {
			insert(node.right, val)
		} else {
			node.right = &BSTNode{val, nil, nil}
		}
	}
}

func contains(node *BSTNode, val Comparable) bool {
	switch {
	case node == nil:
		return false
	case node.val == val:
		return true
	case val.Compare(node.val) < 0:
		return contains(node.left, val)
	case val.Compare(node.val) >= 0:
		return contains(node.right, val)
	}
	return false
}

func remove(node *BSTNode, val Comparable) (*BSTNode, bool) {
	if node != nil {
		switch val.Compare(node.val) {
		case 0:
			// we have to delete this current node
			appendToNode(node.right, node.left)
			return node.right, true
		case -1:
			var ok bool
			var nl *BSTNode
			if nl, ok = remove(node.left, val); ok {
				node.left = nl
			}
			return node, ok
		case 1:
			var ok bool
			var nr *BSTNode
			if nr, ok = remove(node.right, val); ok {
				node.right = nr
			}
			return node, ok
		}
	}
	return node, false
}

// This function appends sub node to a parent node preserving BST properties
func appendToNode(node *BSTNode, sub *BSTNode) {
	// if either of them are nil we do not have to do anything
	if node == nil || sub == nil {
		return
	}
	switch {
	case sub.val.Compare(node.val) < 0:
		if node.left == nil {
			node.left = sub
		} else {
			appendToNode(node.left, sub)
		}
	case sub.val.Compare(node.val) >= 0:
		if node.right == nil {
			node.right = sub
		} else {
			appendToNode(node.right, sub)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func height(node *BSTNode) int {
	if node == nil {
		return 0
	} else {
		return 1 + max(height(node.left), height(node.right))
	}
}

func appendElemsToSlice(node *BSTNode, res []Comparable) []Comparable {
	if node != nil {
		res = append(res, node.val)
		res = appendElemsToSlice(node.left, res)
		res = appendElemsToSlice(node.right, res)
		return res
	}
	return res
}
