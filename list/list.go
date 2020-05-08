package list

import (
	"fmt"
)

// List type
// type List []*g.Game
type List []Item

// Item inf
type Item interface {
	Print()
	Discount(float64)
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
	fmt.Println("Print ....")
	for _, v := range l {
		v.Print()
	}
}

// Discount func
func (l List) Discount(discount float64) {
	fmt.Printf("Discounting by %v  \n.", discount) // .*bookgame.Game game? true (&{gta 60})

	// "it" here is Item
	for _, it := range l {
		it.Discount(discount)
	}
}

// // Discount func
// func (l List) Discount(discount float64) {
// 	fmt.Printf("Discounting by %v  \n.", discount) // .*bookgame.Game game? true (&{gta 60})

// 	for _, it := range l {
// 		// similar to instanceOf in java
// 		// g, ok := it.(*g.Game)
// 		// fmt.Printf("%T game? %v (%v) \n.", it, ok, g) // .*bookgame.Game game? true (&{gta 60})

// 		// if ok {
// 		// 	// fmt.Printf("Discounting %T game? %v (%v) \n.", it, ok, g) // .*bookgame.Game game? true (&{gta 60})
// 		// 	g.Discount(discount)
// 		// }

// 		// check if item has Discount method
// 		// g, ok := it.(interface{ Discount(discount float64) })
// 		g, ok := it.(interface{ Discount(float64) })
// 		if !ok {
// 			continue
// 		}
// 		g.Discount(discount)

// 	}
// }

// // Discount func
// func (l List) Discount(discount float64) {
// 	fmt.Printf("Discounting by %v  \n.", discount) // .*bookgame.Game game? true (&{gta 60})

// 	// can declare interface a func level
// 	type discounter interface {
// 		Discount(float64)
// 	}
// 	// "it" here is Item
// 	for _, it := range l {
// 		// here the type assertion doesn't extract the dynamic value.
// 		// instead it converts the Item to a discounter interface value,
// 		// but only if the value has a discount method
// 		// "it" is now discounter
// 		if it, ok := it.(discounter); ok {
// 			fmt.Printf("discounting (%v) \n.", it)
// 			it.Discount(discount)
// 		}

// 	}
// }
