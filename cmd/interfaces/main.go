package main

import (
	"goMethodsOOP/bookgame"
	l "goMethodsOOP/list"
)

func main() {
	mobyDick := bookgame.Book{Title: "Moby Dick", Price: 20}
	doom := bookgame.Game{Title: "Doom", Price: 40}
	gta := bookgame.Game{Title: "gta", Price: 60}
	risk := bookgame.Puzzle{Title: "risk", Price: 10}
	yoda := bookgame.Toy{Title: "yoda", Price: 310}

	var store l.List
	store = append(store, &doom, &gta, &mobyDick, &risk, &yoda)

	store.Print()

	// fmt.Println(store[0] == &doom)
	// fmt.Println(store[3] == &risk)

	// var p l.Printer

	// p = &doom
	// p = &risk
	// fmt.Println(p)

	store.Discount(.5)
	store.Print()

}
