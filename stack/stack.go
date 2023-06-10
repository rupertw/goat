package stack

type Stack[E any] interface {
	// Push Pushes an item onto the top of this stack and returns the item to be pushed onto this stack.
	Push(e E) E

	// Pop Removes the object at the top of this stack and returns that object as the value of this function.
	Pop() (E, bool)

	// Peek Looks at the object at the top of this stack without removing it from the stack.
	Peek() (E, bool)

	// Len Returns the number of items in this stack.
	Len() int

	// IsEmpty Tests if this stack has no items.
	IsEmpty() bool
}
