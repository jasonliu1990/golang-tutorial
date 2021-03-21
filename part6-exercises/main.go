package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("======ex1======")
	fmt.Println(ex1Foo())
	fmt.Println(ex1Bar())
	fmt.Println("======ex2======")
	x2List := []int{1,2,3,4,5,6,7,8,9,10}
	x2 := ex2Foo(x2List...)
	y2 := ex2Bar(x2List)
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println("======ex3======")
	ex3()
	fmt.Println("======ex4======")
	ex4()
	fmt.Println("======ex5======")
	ex5()
	fmt.Println("======ex6======")
	ex6()
	fmt.Println("======ex7======")
	ex7()
	fmt.Println("======ex8======")
	ex8()
	fmt.Println("======ex9======")
	ex9()
	fmt.Println("======ex10======")
	ex10()
}

// ex1
// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results
func ex1Foo() int{
	return 42
}
func ex1Bar() (int, string){
	return 42, "hello"
}

// ex2
// a. create a func with the identifier foo that
// 		1. takes in a variadic parameter of type int
// 		2. pass in a value of type []int into your func (unfurl the []int)
// 		3. returns the sum of all values of type int passed in
// b. create a func with the identifier bar that
// 		1. takes in a parameter of type []int
// 		2. returns the sum of all values of type int passed in
func ex2Foo(x ...int) int{
	ans := 0
	for _, v := range x{
		ans += v
	}
	return ans
}

func ex2Bar(x []int) int{
	ans := 0
	for _, v := range x{
		ans += v
	}
	return ans
}

// ex3
// Use the “defer” keyword to show that a deferred func runs after the
// func containing it exits.
func ex3Foo(){
	fmt.Println("hello")
}
func ex3Bar(){
	fmt.Println("world")
}
func ex3(){
	defer ex3Foo()
	ex3Bar()
}

// ex4
// a. Create a user defined struct with
//		1. the identifier “person”
//		2. the fields:
//			first
//			last
//			age
// b. attach a method to type person with
//		1. the identifier “speak”
//		2. the method should have the person say their name and age
// c. create a value of type person
// d. call the method from the value of type person

type person struct {
	first string
	last string
	age int
}

func (p person)speak(){
	fmt.Println("I'm", p.first, p.last, ",", p.age, "years old")
}

func ex4(){
	p := person{
		first: "jaz",
		last: "liu",
		age: 30,
	}
	p.speak()
}

// ex5
// a. create a type SQUARE
// b. create a type CIRCLE
// c. attach a method to each that calculates AREA and returns it
// 		1. circle area= π r 2
// 		2. square area = L * W

type square struct {
	length float64
}
type circle struct { 
	radius float64
}
func (s square) area() float64{
	return s.length * s.length
}
func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}
// d. create a type SHAPE that defines an interface as anything
//    that has the AREA method
// e. create a func INFO which takes type shape and then prints the area
type shape interface {
	area() float64
}
func info(s shape){
	fmt.Println(s.area())
}
// f. create a value of type square
// g. create a value of type circle
// h. use func info to print the area of square
// i. use func info to print the area of circle
func ex5(){
	s := square{
		length: 5,
	}
	c := circle{
		radius: 3,
	}
	info(s)
	info(c)
}
// ex6
// Build and use an anonymous func
func ex6(){
	func() {
		fmt.Println("hihi")
	}()
}
// ex7
// Assign a func to a variable, then call that func
func ex7(){
	x := func() {
		fmt.Println("hihihi")
	}
	x()
}
// ex8
// a. Create a func which returns a func
// b. assign the returned func to a variable
// c. call the returned func
func ex8Foo() func() int{
	return func() int {
		return 42
	}
}
func ex8(){
	x := ex8Foo()
	fmt.Println(x())
}
// ex9
// pass a func into a func as an argument
func ex9(){
	g := func(name string) string{
		return fmt.Sprintf("Hi, %s", name)
	}
	fmt.Println(g("Jaz"))
	x := ex9Foo(g, "Jaz")
	fmt.Println(x)
}
func ex9Foo(f func(s string) string, name string) string{
	finalString := f(name)
	return fmt.Sprintf("%s. how's going", finalString)
}

// ex10
// Closure is when we have “enclosed” the scope of a variable in some code block.
// For this hands-on exercise, create a func which “encloses” the scope of a variable:
func ex10(){
	a := ex10Foo()
	fmt.Println(a())
	fmt.Println(a())
	b := ex10Foo()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}
func ex10Foo() func() int{
	x := 0
	return func() int{
		x++
		return x
	}
}
