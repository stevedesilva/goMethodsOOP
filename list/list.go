package list

import (
	"fmt"
	g "goMethodsOOP/bookgame"
)

// List type 
type List []*g.Game

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
