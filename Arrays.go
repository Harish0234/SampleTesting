package main

import "fmt"

// func changeLocal(num [8]int) {
// 	num[0] = 5
// 	fmt.Println("Inside Function", num)
// }

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// var a [3]int // int array with the length 3
	// a[0] = 12    // array index starts at zero('0')
	// a[1] = 10
	// a[2] = 13
	// fmt.Println(a)

	//********Array with short hand declaration****************//
	// a := [3]int{10, 20, 30}
	// fmt.Println("value of a", a)
	//*********Without declaring the length********************//

	// a := [...]int{15, 20, 30}
	// fmt.Println(a)

	//****************Arrays are value Type***********************//

	// a := [...]string{"usa", "india", "china", "Russia"}
	// b := a // a copy of a is assigned to b
	// b[0] = "singapore"
	// fmt.Println("Value of a ", a)
	// fmt.Println("value of b", b)

	//**************Passing the array to a function****************//

	// num := [...]int{15, 20, 5, 40, 45, 85, 25, 785}
	// fmt.Println("Before passing the values to inside function", num)
	// changeLocal(num)
	// fmt.Println("After passing the value", num)

	//***************Length of an array************************//
	// a := [...]int{1, 2, 3, 4, 5}
	// fmt.Println("Length of an array is", len(a))

	//*********************Iterating over elements in the array*********************//
	// a := [...]float64{1, 2, 3, 4, 5, 6, 7, 8}

	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%d th element of a is %.2f\n", i, a[i])
	// }

	// sum := float64(0)
	// for _, v := range a {
	// 	fmt.Printf("the element of a is %.2f\n", v)
	// 	sum += v
	// }
	// fmt.Println("\nsum of all elements of a", sum)
	//////////********************************************Multidimensional array*********************************/////////

	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	printArray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(b)
}
