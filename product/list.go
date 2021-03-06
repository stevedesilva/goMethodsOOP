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
	var str strings.Builder
	for _, p := range l {
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

// Discount func
func (l List) Discount(discount float64) {
	fmt.Printf("Discounting by %v  \n", discount)

	// "it" here is Item
	for _, p := range l {
		p.Discount(discount)
	}
}
