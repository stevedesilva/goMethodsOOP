package list

import (
	"fmt"
)

// List type
// type List []*g.Game
type List []Printer

// Printer inf
type Printer interface {
	Print()
}

// Listprint functionality
func (l List) Listprint() {

	if len(l) == 0 {
		fmt.Println("No games available to print.")
		return
	}
	// - print
	for _, v := range l {
		v.Print()
	}
}

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
