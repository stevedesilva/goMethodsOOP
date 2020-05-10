package main

// "fmt"
import (
	"fmt"
	p "goMethodsOOP/product"
	"sort"
)

func main() {
	l := p.List{
		&p.Product{Title: "Moby Dick", Price: 20},
		&p.Product{Title: "Astrix", Price: 25, Released: p.ToTimestamp(234244979)},
		&p.Product{Title: "Tintin", Price: 40, Released: p.ToTimestamp("123456959")},
	}
	l.Discount(.5)
	// var pocket m.Money = 10
	// fmt.Println("I have ", pocket)

	sort.Sort(l)
	sort.Sort(sort.Reverse(l))
	fmt.Print(l)
	
	sort.Sort(p.ByReleaseDate(l))
	fmt.Print(l)

	sort.Sort(sort.Reverse(p.ByReleaseDate(l)))
	fmt.Print(l)
	

}
