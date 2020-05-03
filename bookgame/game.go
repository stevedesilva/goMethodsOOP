package bookgame

import "fmt"

import m "goMethodsOOP/money"

// Game struct
type Game struct {
	Title string
	Price m.Money
}

// Print func
func (g *Game) Print() {
	fmt.Printf("%-15s: %s\n", g.Title, g.Price.ToString())
}

// Discount price
func (g *Game) Discount(ratio float64) {
	g.Price *= m.Money(1 - ratio)
}