package goadt

type binNode struct {
	val int // TODO: Change this to comparable
	left, right *binNode
}

type BinTree struct {
	head *binNode
}

func insertToBinNode(node *binNode, n int) {
	switch {
	case node == nil: node = &binNode{n, nil, nil}
	case n < node.val: insertToBinNode(node.left, n)
	default: insertToBinNode(node.right, n)
	}
}

func searchNode(node *binNode, n int) bool {
	switch {
	case node == nil: return false
	case n < node.val: return searchNode(node.left, n)
	case n > node.val: return searchNode(node.right, n)
	default: return true
	}
}

func (t *BinTree) Insert(n int) {
	if t.head == nil {
		t.head = &binNode{n, nil, nil}
	} else {
		insertToBinNode(t.head, n)
	}
}

func (t *BinTree) Contains(n int) bool {
	return searchNode(t.head, n)
}
