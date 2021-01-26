package main

import "fmt"

var y = 42
var z = "hello"
var a = `I said, "hello"`

func main()  {
	fmt.Println(y)
	// type
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
