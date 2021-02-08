package main

import "fmt"

// var keyword
// declare a variable outside of a function body
var y = 43
// declare a variable with the identifier z is of type int
// assigns the zero value of type int to z
var z int
func main()  {
	// short declaration operator
	// declare a variable and assign a value
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo(){
	fmt.Println("again:", y)
}
