package main

import "fmt"

func main() {
	// emptyOne()

	emptyTwo()

	// var many []interface{}
	// many = nums
}

func emptyTwo() {
	nums := []int{1, 2, 3}
	var any interface{}
	// CAN assign anything to interface{} type
	any = nums
	fmt.Printf("1 Dynamic type %T \tDynamic value %[1]v\n", any)

	// CANNOT assign interface{} to []int type
	//nums = any

	// you cannot use any as an int slice. any is an empty interface value , it isn't an int slice
	// _ = len(any)
	_ = len(any.([]int))

	var many []interface{}
	// many = nums

	/*
		'many' is a slice of interface{} values
		So 'many' cannot store any type of value (i.e. since it's type is slice of values).
		Only an interface{} value can store any type of value
	*/
	for _, v := range nums {
		many = append(many, v)
	}
	fmt.Println(many)

}

func emptyOne() {
	var any interface{}
	// Dynamic type <nil>    Dynamic value <nil>
	fmt.Printf("1 Dynamic type %T \tDynamic value %[1]v\n", any)

	// you can store any value
	any = []int{1, 2, 3}
	any = map[int]bool{1: true, 2: false}
	any = "hello"
	any = 1

	// you cannot directly use the dynamic value of an empty interface value (2 * any)
	// Dynamic type []int    Dynamic value [1 2 3]
	any = 2 * any.(int)
	fmt.Println(any)
}
