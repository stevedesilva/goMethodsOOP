package bookgame

import "fmt"

// Huge struct
type Huge struct {
	games [10e6]Game
}

// AddressVal func
// copies the a value of huge struct each time it is called
// time go run cmd/bookgame/main.go
// 0xc00e60c000
// 0xc01cb02000
// 0xc02afe4000
// 0xc0394c6000
// 0xc0479a8000
// 0xc055eba000
// 0xc00e60c000
// 0xc01cb02000
// 0xc06439c000
// 0xc0394c6000
// copies the a value of huge struct each time it is called
// time go run cmd/bookgame/main.go
// real    0m1.740s
// user    0m2.044s
// sys     0m0.951s
func (h Huge) AddressVal() {
	fmt.Printf("%p\n", &h)
}

// AddressPtr func
// only copies pointer value
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
// 0xc0000a6000
//
// real    0m1.170s
// user    0m1.179s
// sys     0m0.219s
func (h *Huge) AddressPtr() {
	fmt.Printf("%p\n", h)
}
