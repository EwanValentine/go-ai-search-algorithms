package constraints

// Key is a generic type constraint, as we want to accept numbers
// or strings as a Key to our data structures
type Key interface {
	~string | ~int
}
