package goadt

type node struct {
	val interface{}
	next *node
}

type Stack struct {
	top *node
	size int
}

func (s *Stack) Push(v interface{}) *Stack {
	s.top = &node{v,s.top} 
	s.size++
	return s
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.top == nil {
		return nil, false
	}
	return s.top.val, true
}

func (s *Stack) Pop() (v interface{}, m bool) {
	v, m = nil, false
	if s.top != nil {
		v, s.top = s.top.val, s.top.next
		m = true
		s.size--
	} 
	return
}

func (s *Stack) Size() int {
	return s.size
}
