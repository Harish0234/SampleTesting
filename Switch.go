package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	// finger := 10
	// fmt.Printf("finger %d is", finger)
	// switch finger {
	// case 1:
	// 	fmt.Println("Thumb")
	// case 2:
	// 	fmt.Println("index")
	// case 3:
	// 	fmt.Println("middle")
	// case 4:
	// 	fmt.Println("Ring")
	// // case 4: // duplicates cases are not allowed in the switch such compiler throws the error
	// // 	fmt.Println("another ring")
	// case 5:
	// 	fmt.Println("pinky")
	// default:
	// 	{ // default case will excute if we don't have other case to excute
	// 		fmt.Println("Incorrect finger number")
	// 	}
	// }
	//**************Multiple expression in Switch********************************//

	// letter := "i"
	// fmt.Printf("Letter %s is a", letter)
	// switch letter {
	// case "a", "e", "i", "o", "u":
	// 	fmt.Println("vowel")
	// default:
	// 	fmt.Println("not vowel")
	// }

	//*************Expression Less Switch********************//

	// num := 75

	// switch { // Switch statement without expression
	// case num > 0 && num < 50:
	// 	fmt.Printf("%d is greater than 0 and less than 50", num)
	// case num > 51 && num < 70:
	// 	fmt.Printf("%d is greater than 51 and less than 70", num)
	// case num > 71 && num < 76:
	// 	fmt.Printf("%d is greater than 71 and less than 76", num)
	// }

	//*******************************Fall Throught***********************************//

	switch num := number(); {
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 75:
		fmt.Printf("%d is lesser than 75\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)

	}

}
