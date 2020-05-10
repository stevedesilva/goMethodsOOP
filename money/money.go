package money

import "fmt"

// A type and its methods should be in the same package

// Money Type
type Money float64

// ToString func
func (m Money) String() string {
	return fmt.Sprintf("$%.2f", m)
}
