package main

import bg "goMethodsOOP/bookgame"
// https://circleci.com/gh/stevedesilva/workflows/goMethodsOOP
func main() {
	mobyDick := bg.Book{
		Title: "Moby Dick",
		Price: 10,
	}

	doom := bg.Game{
		Title: "Doom",
		Price: 40,
	}

	gta := bg.Game{
		Title: "gta",
		Price: 60,
	}

	mobyDick.Print()
	doom.Print()
	gta.Print()
}
