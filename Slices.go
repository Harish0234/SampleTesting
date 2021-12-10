package main

import (
	"fmt"
)

//*****Syntax of Slice []T(Type)***********//

// func substractOne(number []int) {
// 	for i := range number {
// 		number[i] -= 2
// 	}
// }

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	fmt.Println("Countries which are in the string", countries)
	neededCountries := countries[:len(countries)-2]
	fmt.Println("needed countries", neededCountries)
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	// a := [...]int{15, 20, 25, 4, 54, 45, 78}
	// var b []int = a[1:4] // creating a slice to print the value for 1 to 4
	// fmt.Print(b)

	// b := []int{1, 2, 3, 4} // declaring inside the slice
	// fmt.Print(b)

	// darr := [...]int{1, 2, 3, 4, 5, 6, 8, 5, 1, 2, 5, 6, 2, 6}
	// dslice := darr[2:9] //Modifying a slice
	// fmt.Println("Array before", darr)

	// for i := range dslice {
	// 	dslice[i]++
	// }
	// fmt.Println("Array after", darr)
	//***************slices share the same underlying array, the changes that each one makes will be reflected in the array*****//

	// numa := [...]int{10, 20, 30, 41}
	// nums1 := numa[:] // Slice which contains all elements
	// nums2 := numa[:]
	// fmt.Println("Array before changing 1", numa)
	// nums1[0] = 100
	// fmt.Println("After changing the nums1", numa)
	// nums2[0] = 101
	// fmt.Println("After changing the nums2", numa)

	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6

	//**************re-sliced upto its capacity************************//

	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	// fruitslice = fruitslice[:cap(fruitslice)]                                        //re-slicing furitslice till its capacity
	// fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	//**********************creating a slice using make****************************//

	// i := make([]int, 5, 5)
	// fmt.Println(i)
	//******************Appending to a slice**********************//

	// cars := []string{"suzuki", "Honda", "toyota"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	// cars = append(cars, "Mazada")
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))

	//******************Passing a slice to function**********************//
	//****Syntax*****//
	// type slice struct {
	// 	Length        int
	// 	Capacity      int
	// 	ZerothElement *byte
	// nos := []int{8, 7, 6}
	// fmt.Println("slice before passing the function call", nos)
	// substractOne(nos)
	// fmt.Println("slice after passing the function call", nos)

	//**********************Multi-dimensional Slice*********************************//

	// Data := [][]string{
	// 	{"c", "c++"},
	// 	{"Java", "Dotnet"},
	// 	{"Springboot"},
	// 	{"go", "rust"},
	// }

	// for _, v1 := range Data {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s ", v2)
	// 	}
	// 	println("\n")
	// }

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

}
