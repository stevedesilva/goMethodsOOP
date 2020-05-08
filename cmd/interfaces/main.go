package main

import (
	"goMethodsOOP/bookgame"
	l "goMethodsOOP/list"
)

func main() {
	store := l.List{
		&bookgame.Book{Title: "Moby Dick", Price: 20},
		&bookgame.Book{Title: "Astrix", Price: 40, Published: 234244979},
		&bookgame.Book{Title: "Tintin", Price: 40, Published: "123456959"},
		&bookgame.Game{Title: "Doom", Price: 40},
		&bookgame.Game{Title: "gta", Price: 60},
		&bookgame.Puzzle{Title: "risk", Price: 10},
		&bookgame.Toy{Title: "yoda", Price: 310},
	}

	store.Discount(.5)
	store.Print()
}
