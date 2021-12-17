//*******Synta of a map  **make(map[type of key]type of value)** *************//

package main

import "fmt"

func main() {

	// employeSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// 	"mike":  9000,
	// }
	// employee := "jamie"
	// salary := employeSalary[employee]
	// fmt.Println("salary of", employee, "is", salary)

	//*we want to know whether an key is present in the employeeSalary map.*//
	//**Syntax**//
	//value, ok := map[key]//
	// employeeSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// }
	// newEmp := "joe"
	// value, ok := employeeSalary[newEmp]
	// if ok == true {
	// 	fmt.Println("Salary of", newEmp, "is", value)
	// 	return
	// }
	// fmt.Println(newEmp, "not found")

	//***Iterate over all elements in a map**///
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	// fmt.Println("Contents in the map")
	// for key, value := range employeeSalary {
	// 	fmt.Printf("employeeSalary[%s] = %d\n", key, value)

	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)

}
