package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (lst *List[T]) Push(val T) {
	lst.next = &List[T]{nil, val}
}

func main() {
}
