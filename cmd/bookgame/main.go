package main

import bg "goMethodsOOP/bookgame"

// https://circleci.com/gh/stevedesilva/workflows/goMethodsOOP
func main() {

	mobyDick := &bg.Game{Product: bg.Product{Title: "mobyDick", Price: 10}}
	doom := &bg.Game{Product: bg.Product{Title: "Doom", Price: 10}}
	gta := &bg.Game{Product: bg.Product{Title: "GTA", Price: 90}}

	mobyDick.Print()
	// behind the scenes = (&gta).Discount(0.5)
	gta.Discount(0.5)
	gta.Print()

	doom.Print()
	// behind the scenes =  bg.Game.Print(doom)

	huge := bg.Huge{}

	// for i := 0; i < 10; i++ {
	// 	huge.AddressVal()
	// }

	for i := 0; i < 10; i++ {
		huge.AddressPtr()
	}

}
