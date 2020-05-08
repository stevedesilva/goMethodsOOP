package main

import (
	"fmt"
	bg "goMethodsOOP/bookgame"
	l "goMethodsOOP/list"
)

func main() {
	store := l.List{
		&bg.Book{Product: bg.Product{Title: "Moby Dick", Price: 20}, Published: nil},
		&bg.Book{Product: bg.Product{Title: "Astrix", Price: 25}, Published: 234244979},
		&bg.Book{Product: bg.Product{Title: "Tintin", Price: 40}, Published: "123456959"},
		&bg.Game{Product: bg.Product{Title: "Doom", Price: 10}},
		&bg.Game{Product: bg.Product{Title: "GTA", Price: 90}},
		&bg.Puzzle{Product: bg.Product{Title: "TETRIS", Price: 30}},
		&bg.Puzzle{Product: bg.Product{Title: "JENGA", Price: 30}},
		&bg.Toy{Product: bg.Product{Title: "DARTH VADER", Price: 130}},
	}
	store.Print()
	store.Discount(.5)
	store.Print()

	toy := &bg.Toy{Product: bg.Product{Title: "YODA", Price: 130}}
	// &bookgame.Toy{Product:bookgame.Product{Title:"YODA", Price:130}}
	fmt.Printf("%#v \n", toy) // print field & package & values names
	//&{{YODA 130}}
	fmt.Printf("%v \n", toy) // only print values names
	// &{Product:{Title:YODA Price:130}}
	fmt.Printf("%+v \n", toy) // print field & value names

	book := &bg.Book{Product: bg.Product{Title: "Astrix", Price: 25}, Published: 234244979}
	fmt.Printf("%#v \n", book) // print field & package & values names
	fmt.Printf("%v \n", book)  // only print values names
	fmt.Printf("%+v \n", book) // print field & value names
}
