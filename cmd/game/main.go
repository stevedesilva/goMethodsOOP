package main

import (
	g "goMethodsOOP/bookgame"
	l "goMethodsOOP/list"
)

// https://circleci.com/gh/stevedesilva/workflows/goMethodsOOP
func main() {

	doom := g.Game{Title: "Doom", Price: 40}
	gta := g.Game{Title: "gta", Price: 60}

	// - create array game pointer

	// Alternative
	// items := []*g.Game(nil)
	// or	
	// var items []*g.Game

	// // - append to item
	// items = append(items, &doom, &gta)

	items := []*g.Game{&doom, &gta}

	// Cast to List (i.e.type List []*g.Game)
	// Note: You can convert types that have identical underlying type
	// List type and items share same underlying type []*g.Game
	// casting to List decorates items with List methods (e.g. Listprint)
	// could use new
	my := l.List(items)

	// you can attach methods to a nil valiue
	// my = nil

	// l.Listprint(items)
	// 
	my.Listprint()
}
