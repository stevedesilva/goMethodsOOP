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
func (p *Product) Print() {
	fmt.Printf("%s: %s (%s)\n", p.Title, p.Price.ToString(), p.Released.String())
}

// Discount price
func (p *Product) Discount(ratio float64) {
	p.Price *= m.Money(1 - ratio)
}
