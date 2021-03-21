package main

import "fmt"

func main(){
	fmt.Println("======ex1======")
	ex1()
	fmt.Println("======ex1======")
	ex2()
}

// ex1
// Create a value and assign it to a variable.
// Print the address of that value.
func ex1(){
	a := 42
	fmt.Println(&a)
}

// ex2
// a. create a person struct
// b. create a func called “changeMe” with a *person as a parameter
//		1. change the value stored at the *person address
//			important
//				to dereference a struct, use (*value).field
//					- p1.first
//					- (*p1).first
//				“As an exception, if the type of x is a named pointer type
//				and (*x).f is a valid selector expression denoting a field
//				(but not a method), x.f is shorthand for (*x).f.”
// c. create a value of type person
//		print out the value
// d. in func main
//		call “changeMe”
//		print out the value
type person struct {
	name string
}
func changeMe(p *person)  {
	p.name = "James"
}
func ex2(){
	p1 := person{name: "Jaz"}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
