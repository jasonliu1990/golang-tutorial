package main

import "fmt"

var y string
var z int

func main()  {
	// declare a variable to be of a certain type
	// and then assign a value of that type to the variable
	fmt.Println("printing the value of", y, "end")
	fmt.Printf("%T\n", y)

	y = "Hello"

	fmt.Println("printing the value of", y, "end")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
