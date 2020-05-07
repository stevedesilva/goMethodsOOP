package bookgame

import (
	"fmt"

	m "goMethodsOOP/money"
)

// Toy struct
type Toy struct {
	Title string
	Price m.Money
}

// Print func
func (g *Toy) Print() {
	fmt.Printf("%-15s: %s\n", g.Title, g.Price.ToString())
}

// // Discount price
// func (g *Toy) Discount(ratio float64) {
// 	g.Price *= m.Money(1 - ratio)
// }
