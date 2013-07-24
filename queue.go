package goadt

type Queue struct {
	head *node
	size int
}

func (q *Queue) Put(v interface{}) {
	if q.head == nil {
		q.head = &node{v, nil} 
	} else {
		n := q.head
		for ; n.next != nil; {
			n = n.next
		}
		n.next = &node{v, nil}
	} 
	q.size++
}

func (q *Queue) Remove() (v interface{}) {
	if q.head != nil {
		v, q.head = q.head.val, q.head.next
		q.size--
	}
	return v
}

func (q *Queue) Empty() bool {
	return q.head == nil
}

func (q *Queue) Size() int {
	if q == nil {
		return -1
	}
	return q.size
}
