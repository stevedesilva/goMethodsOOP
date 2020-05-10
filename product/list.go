package product

import (
	"fmt"
	"strings"
)

// List type
type List []*Product

// Print func
func (l List) String() string {

	if len(l) == 0 {
		return "Sorry, We're waiting for a delivery.\n"
	}

	// - print
	}
	return str.String()
}

// Discount func
func (l List) Discount(discount float64) {
	fmt.Printf("Discounting by %v  \n.", discount)

	// "it" here is Item
	for _, p := range l {
		p.Discount(discount)
	}
}
