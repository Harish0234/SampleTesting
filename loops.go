package main

import "fmt"

func main() {
	//num := 100
	// if num%2 == 0 {
	// 	fmt.Print("The Number", num, "is even")
	// 	return
	// }
	// fmt.Print("the number", num, "is odd")
	//If Else condition
	// if num/10 == 10 {
	// 	fmt.Println("The given number is even", num)
	// } else {
	// 	fmt.Println("Given number is odd", num)
	// }
	// **else if**
	// if num <= 50 {
	// 	fmt.Println(num, "Is less then or equal to 50")
	// } else if num >= 51 && num <= 99 {
	// 	fmt.Println(num, "lies between 51 to 99")
	// } else {
	// 	fmt.Println(num, "is equal to 100")
	// }

	//*assignment-statement*

	if num := 10; num%2 == 0 {
		fmt.Println("Given number is even")
	} else {
		fmt.Println("Given number is odd")
	}

}
