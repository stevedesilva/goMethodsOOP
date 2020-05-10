package main

// "fmt"
import (
	"fmt"
	// p "goMethodsOOP/product"
	m "goMethodsOOP/money"
)

func main() {
	// l := p.List{
	// 	&p.Product{Title: "Moby Dick", Price: 20},
	// 	&p.Product{Title: "Astrix", Price: 25, Released: p.ToTimestamp(234244979)},
	// 	&p.Product{Title: "Tintin", Price: 40, Released: p.ToTimestamp("123456959")},
	// }
	// l.Discount(.5)
	// fmt.Print(l)
		var pocket m.Money = 10
		fmt.Println("I have ",pocket)

}
