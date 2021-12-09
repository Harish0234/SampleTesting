package main

import (
	"fmt"
	"log"
	"mymodule/simpleinterest"
)

var p, r, t = -5000.0, 10.0, 1.0

func main() {
	fmt.Println("main package initialized")
	if p < 0 {
		log.Fatal("principal is less then Zero")
	}
	if r < 0 {
		log.Fatal("rate of interest is less then Zero")
	}
	if t < 0 {
		log.Fatal("Duration is less then Zero")
	}

	// p := 5000.0
	// r := 10.00
	// t := 300.00
	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Final value", si)

}
