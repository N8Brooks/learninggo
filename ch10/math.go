// Currently, does not work for importing as this is a subdirectory.
package ch10

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two integers and returns the sum of them.
// Go to https://www.mathsisfun.com/numbers/addition.html for more information.
func Add[T Number](a, b T) T {
	return a + b
}
