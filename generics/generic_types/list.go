package main

// List represent a singly-linked list that holds
// values of any type
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

}
