package main

import "fmt"

func main() {
	var any interface{}
	// you can store any value
	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "hello"
	any = 1

	// you cannot directly use the dynamic value of an empty interface value (2 * any)

	any = 2 * any.(int)
	fmt.Println(any)

	nums := []int{1, 2, 3}
	any = nums
	fmt.Printf("Dynamic type %T \tDynamic value %[1]v\n",any)
}
