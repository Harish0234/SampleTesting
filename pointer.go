//********Pointer**************//
//***A pointer is a variable that stores the memory address of another variable.**.//

package main

import "fmt"

// func change(val *int) {
// 	*val = 55
// }

// func hello() *int {
// 	i := 5
// 	return &i
// }

// func modify(arr *[3]int) {
// 	(*arr)[0] = 90

func modify(arr *[3]int) { //using this shorthand syntax.
	arr[0] = 90
}

func main() {
	// b := 225
	// var a *int = &b
	// fmt.Printf("Type of a is %T\n", a)
	// fmt.Println("address of b is", a)

	//The zero value of a pointer is nil.

	// a := 25
	// var b *int

	// if b == nil {
	// 	fmt.Println("b is", b)
	// 	b = &a
	// 	fmt.Println("b after initialization is", b)
	// }

	// ** Creating pointers using the new function** //

	// size := new(int)
	// fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	// *size = 85
	// fmt.Println("New size value is", *size)

	//**Dereferencing a pointer**//
	// b := 255
	// a := &b
	// fmt.Println("address of b is", a)
	// fmt.Println("value of b is", *a) //Dereferencing a pointer means accessing the value of the variable to which the pointer points. *a is the syntax to deference a
	// *a++                             //where we change the value in b using the pointer.
	// fmt.Println("value of b is", *a)

	// a := 58
	// fmt.Println("value of a before function call is", a)
	// b := &a
	// change(b)
	// fmt.Println("value of a after function call is", a)

	// d := hello()
	// fmt.Println("Value of d", *d)

	a := [3]int{89, 90, 91}
	modify(&a)
	fmt.Println(a)
}
