package main

// "fmt"
import p "goMethodsOOP/product"

func main() {
	store := p.List{
		&p.Product{Title: "Moby Dick", Price: 20},
		&p.Product{Title: "Astrix", Price: 25, Released: p.ToTimestamp(234244979)},
		&p.Product{Title: "Tintin", Price: 40, Released: p.ToTimestamp("123456959")},
	}
	store.Discount(.5)
	store.Print()

}
