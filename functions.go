package main

import "fmt"

// func calculate(price int, no int) int {
// 	total := price * no
// 	return total
// }

// func main() {
// 	Totalprice := calculate(10, 20)
// 	fmt.Println("TotalValue", Totalprice)
// }

// func rectprops(length, width float32) (float32, float32) {
// 	area := length * width
// 	perimeter := (length + width) / 2
// 	return area, perimeter
// }

// func main() {
// 	area, perimeter := rectprops(15.0, 20.2)
// 	fmt.Println("rectangle", area, perimeter)
// }

func rectprops(length, width float32) (float32, float32) {
	area := length * width
	perimeter := (length + width) / 2
	return area, perimeter
}
func main() {
	area, _ := rectprops(15.00, 20.00)
	fmt.Println("rectangle", area)
}
