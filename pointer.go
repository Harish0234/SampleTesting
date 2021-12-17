//********Pointer**************//
//***A pointer is a variable that stores the memory address of another variable.**.//

package main

import "fmt"

func main() {
	// b := 225
	// var a *int = &b
	// fmt.Printf("Type of a is %T\n", a)
	// fmt.Println("address of b is", a)

	//The zero value of a pointer is nil.

	a := 25
	var b *int

	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}

}
