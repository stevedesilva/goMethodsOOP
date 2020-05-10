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
	for _, v := range l {
		v.Print()
	}
}

// Discount func
func (l List) Discount(discount float64) {
	fmt.Printf("Discounting by %v  \n", discount)

	// "it" here is Item
	for _, p := range l {
		p.Discount(discount)
	}
}

// Len is the number of elements in the collection.
func (l List) Len() int {
	return len(l)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (l List) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}

// Swap swaps the elements with indexes i and j.
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
