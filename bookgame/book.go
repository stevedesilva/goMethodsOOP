package bookgame

import "fmt"

// receiver name should be consistent with previous receiver name b for Book

// Book struct
type Book struct {
	Title string
	Price float64
}

// Print func
func (b *Book) Print() {
	fmt.Printf("%-15s: $%.2f\n", b.Title, b.Price)
}

// Discount price
func (b *Book) Discount(ratio float64) {
	b.Price *= (1 - ratio)
}
