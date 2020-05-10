package bookgame

import (
	"fmt"
	m "goMethodsOOP/money"
)

// Product struct
type Product struct {
	Title string
	Price m.Money
}

// Print func
func (p *Product) Print() {
	fmt.Printf("%-15s: %s\n", p.Title, p.Price)
}

// Discount price
func (p *Product) Discount(ratio float64) {
	p.Price *= m.Money(1 - ratio)
}
