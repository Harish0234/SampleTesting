//**A struct is a user-defined type that represents a collection of fields**//

//Declaring a struct
// type Employee struct {
//     firstName string
//     lastName  string
//     age       int
// }

package main

import "fmt"

// type Employee struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	salary    int
// }

//**Anonymous fields**//
//**It is possible to create structs with fields that contain only a type without the field name. These kinds of fields are called anonymous fields.**//

// type Person struct {
// 	string
// 	int
// }

//**Nested structs**//
// type Address struct {
// 	city  string
// 	state string
// }

// type Person struct {
// 	name    string
// 	age     int
// 	address Address
// }

//**Promoted fields**//

type Address struct {
	city  string
	state string
}
type Person struct {
	name    string
	age     int
	Address // promoted field (the Person struct has an anonymous field Address which is a struct. Now the fields of the Address namely city and state are called promoted fields since they can be accessed as if they are directly declared in the Person struct itself.)
}

func main() {

	// emp1 := Employee{"Harish", "Gowrishetti", 25, 50000000}

	// emp2 := Employee{"Rajesh", "sudam", 27, 50000000}

	// fmt.Println("Employee1", emp1)
	// fmt.Println("Employee1", emp2)

	//** Creating anonymous structs **//

	// emp3 := struct {
	// 	firstName string
	// 	lastName  string
	// 	age       int
	// 	salary    int
	// }{
	// 	firstName: "Andreah",
	// 	lastName:  "Nikola",
	// 	age:       31,
	// 	salary:    5000,
	// }

	// fmt.Println("Employee 3", emp3)

	//**Accessing individual fields of a struct**//
	// emp6 := Employee{
	// 	firstName: "Sam",
	// 	lastName:  "Anderson",
	// 	age:       55,
	// 	salary:    6000,
	// }
	// fmt.Println("First Name:", emp6.firstName)
	// fmt.Println("Last Name:", emp6.lastName)
	// fmt.Println("Age:", emp6.age)
	// fmt.Printf("Salary: $%d\n", emp6.salary)
	// emp6.salary = 6500
	// fmt.Printf("New Salary: $%d", emp6.salary)

	//**Zero value of a struct**//

	// var emp4 Employee //zero valued struct
	// fmt.Println("First Name:", emp4.firstName)
	// fmt.Println("Last Name:", emp4.lastName)
	// fmt.Println("Age:", emp4.age)
	// fmt.Println("Salary:", emp4.salary)

	// p1 := Person{
	// 	string: "naveen",
	// 	int:    50,
	// }
	// fmt.Println(p1.string)
	// fmt.Println(p1.int)

	p := Person{
		name: "Naveen",
		age:  50,
		Address: Address{
			city:  "Chicago",
			state: "Illinois",
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)
	fmt.Println("State:", p.state)

}
