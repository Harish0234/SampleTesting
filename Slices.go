package main

import "fmt"

//*****Syntax of Slice []T(Type)***********//

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

	numa := [...]int{10, 20, 30, 41}
	nums1 := numa[:] // Slice which contains all elements
	nums2 := numa[:]
	fmt.Println("Array before changing 1", numa)
	nums1[0] = 100
	fmt.Println("After changing the nums1", numa)
	nums2[0] = 101
	fmt.Println("After changing the nums2", numa)

}
