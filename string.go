package main

import "fmt"

//*** %X & %x "x" and "X" serve to output a hexadecimal number.
//"x" stands for lower case letters (abcdef) while "X" for capital letters (ABCDEF).//

// func printBytes(s string) {
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x 	", s[i])
// 	}
// }
// func printChar(s string) {
// 	rune := []rune(s)
// 	for i := 0; i < len(rune); i++ {
// 		fmt.Printf("%c 	", rune[i])
// 	}
// }

// func charandbytes(s string) {
// 	for index, rune := range s {
// 		fmt.Printf("%c starts at byte %d\n", rune, index)
// 	}
// }

// func compareStrings(str1 string, str2 string) {
// 	if str1 == str2 {
// 		fmt.Printf("%s and %s are equal\n", str1, str2)
// 		return
// 	}
// 	fmt.Printf("%s and %s are not equal\n", str1, str2)
// }

// func mutate(s string) string {
// 	s[0] = 'a'//any valid unicode character within single quote is a rune
// 	return s
// }

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func main() {
	// String := "HarisH"
	// fmt.Println("String:", String)
	// printBytes(String)
	// fmt.Printf("\n")
	// printChar(String)
	// fmt.Printf("\n\n")
	// name := "Señor"
	// printBytes(name)
	// fmt.Printf("\n")
	// printChar(name)
	// fmt.Printf("\n")

	// name := "Señor"
	// charandbytes(name)

	// byteSlice := []byte{0x48, 0x61, 0x72, 0x69, 0x73, 0x48}
	// str := string(byteSlice)
	// fmt.Println(str)

	// runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	// str1 := string(runeSlice)
	// fmt.Println(str1)

	// word1 := "Señor"
	// fmt.Printf("String: %s\n", word1)
	// fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	// fmt.Printf("Number of bytes: %d\n", len(word1))

	// fmt.Printf("\n")
	// word2 := "Pets"
	// fmt.Printf("String: %s\n", word2)
	// fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	// fmt.Printf("Number of bytes: %d\n", len(word2))

	// string1 := "Go"
	// string2 := "Go"
	// compareStrings(string1, string2)

	// string3 := "hello"
	// string4 := "world"
	// compareStrings(string3, string4)

	// String1 := "Harish"
	// String2 := "is superhero"
	// StringConcat := String1 + " " + String2
	// fmt.Println(StringConcat)
	// stringfinal := fmt.Sprintf("%s,%s", String1, String2)
	// fmt.Println(stringfinal)

	h := "hello"
	fmt.Println(mutate([]rune(h)))

}
