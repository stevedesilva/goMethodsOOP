package bookgame

import (
	"fmt"
	m "goMethodsOOP/money"
)

// receiver name should be consistent with previous receiver name b for Book

// Book struct
type Book struct {
	Title string
	Price m.Money
}

// Print func
func (b *Book) Print() {
	fmt.Printf("%-15s: %s\n", b.Title, b.Price.ToString())
}

// Discount price
func (b *Book) Discount(ratio float64) {
	b.Price *= m.Money(1 - ratio)
}
