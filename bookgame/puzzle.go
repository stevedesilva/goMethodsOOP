package bookgame

import (
	"fmt"
	m "goMethodsOOP/money"
)

// receiver name should be consistent with previous receiver name b for Book

// Puzzle struct
type Puzzle struct {
	Title string
	Price m.Money
}

// Print func
func (p *Puzzle) Print() {
	fmt.Printf("%-15s: %s\n", p.Title, p.Price.ToString())
}
