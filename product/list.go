package product

import (
	"fmt"
)

// List type
type List []*Product

// Print func
func (l List) Print() {

	if len(l) == 0 {
		fmt.Println("No games available to print.")
		return
	}
	// - print
	fmt.Println("Print ....")
	for _, v := range l {
		v.Print()
	}
}

// Discount func
func (l List) Discount(discount float64) {
	fmt.Printf("Discounting by %v  \n.", discount)

	// "it" here is Item
	for _, p := range l {
		p.Discount(discount)
	}
}
