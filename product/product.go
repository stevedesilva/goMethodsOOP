package product

import (
	"fmt"
	m "goMethodsOOP/money"
)

// Product struct
type Product struct {
	Title    string
	Price    m.Money
	Released Timestamp
}

// Print func
func (p *Product) String() string {
	return fmt.Sprintf("%s: %s (%s)", p.Title, p.Price, p.Released)
}

// Discount price
func (p *Product) Discount(ratio float64) {
	p.Price *= m.Money(1 - ratio)
}
