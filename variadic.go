package main

import "fmt"

//*******Syntax**************//
// func variadic (a int, b... int){
// variadic function is a function that accepts a variable number of arguments
//Only the last parameter of a function can be variadic. We will learn why this is the case in the next section of this tutorial
// }

func find(a int, b ...int) {

	fmt.Println("types in the", b)

}
