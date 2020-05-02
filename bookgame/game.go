package bookgame

import "fmt"

// Game struct
type Game struct {
	Title string
	Price float64
}

// Print func
func (g *Game) Print() {
	fmt.Printf("%-15s: $%.2f\n", g.Title, g.Price)
}

// Discount price
func (g *Game) Discount(ratio float64) {
	g.Price *= (1 - ratio)
}