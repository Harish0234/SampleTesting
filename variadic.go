package main

import "fmt"

//*******Syntax**************//
// func variadic (a int, b... int){
// variadic function is a function that accepts a variable number of arguments
//Only the last parameter of a function can be variadic. We will learn why this is the case in the next section of this tutorial
// }

// func find(num int, nums ...int) {
// 	fmt.Printf("type of nums is %T\n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Println(num, "found at index", i, "in", nums)
// 			found = true
// 		}
// 		if !found {
// 			fmt.Println(num, "not found", i)
// 		}
// 	}
// 	fmt.Printf("\n")
// }

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "Playground")
	fmt.Println(s)
}

func main() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	// find(10, 25, 35, 45, 55, 11, 10)
	// find(45, 55, 75, 85, 95, 45, 25)
	// find(15, 78, 25, 45, 18, 54, 284, 14)

}
