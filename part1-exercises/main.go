package main


import (
	"fmt"
)
// ex1
//1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with
//   the IDENTIFIERS “x” and “y” and “z”
//   a. 42; b. “James Bond”; c. true
//2. Now print the values stored in those variables using
//   a single print statement
//   multiple print statements
func ex1(){
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)
}
// ex2
//1. Use var to DECLARE three VARIABLES. The variables should have package
//   level scope. Do not assign VALUES to the variables. Use the following
//   IDENTIFIERS for the variables and make sure the variables are of the
//   following TYPE (meaning they can store VALUES of that TYPE).
//   a. identifier “x” type int
//   b. identifier “y” type string
//   c. identifier “z” type bool
//2. in func main
//   a. print out the values for each identifier
//   b. The compiler assigned values to the variables. What are these values
//      called? <- the zero value
var x2 int
var y2 string
var z2 bool
func ex2(){
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)
}
//ex3
//Using the code from the previous exercise,
//1. At the package level scope, assign the following values to the three
//   variables
//   a. for x assign 42
//   b. for y assign “James Bond”
//   c. for z assign true
//2. in func main
//   a. use fmt.Sprintf to print all of the VALUES to one single string.
//      ASSIGN the returned value of TYPE string using the short declaration
//      operator to a VARIABLE with the IDENTIFIER “s”
//   b. print out the value stored by variable “s”
var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true
func ex3(){
	s := fmt.Sprintf("%v %v %v", x3, y3, z3)
	fmt.Println(s)
}
//ex4
//1. Create your own type. Have the underlying type be an int.
//2. create a VARIABLE of your new TYPE with the IDENTIFIER “x”
//   using the “VAR” keyword
//3. in func main
//   a. print out the value of the variable “x”
//   b. print out the type of the variable “x”
//   c. assign 42 to the VARIABLE “x” using the “=” OPERATOR
//   d. print out the value of the variable “x”

type mytype int
var x4 mytype
func ex4(){
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
}

//ex5
//1. Building on the code from the previous example
//   at the package level scope, using the “var” keyword,
//   create a VARIABLE with the IDENTIFIER “y”.
//   The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
//2. in func main
//    a. this should already be done
//		i. print out the value of the variable “x”
//		ii. print out the type of the variable “x”
//		iii. assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
//		iv. print out the value of the variable “x”
//    b. now do this
//		i. now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the
//	       UNDERLYING TYPE
//			1. then use the “=” operator to ASSIGN that value to “y”
//			2. print out the value stored in “y”
//			3. print out the type of “y”
var x5 mytype
var y5 int
func ex5(){
	x5 = 42
	y5 = int(x5)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}



func main() {
	fmt.Println("====== ex1 ======")
	ex1()
	fmt.Println("====== ex2 ======")
	ex2()
	fmt.Println("====== ex3 ======")
	ex3()
	fmt.Println("====== ex4 ======")
	ex4()
	fmt.Println("====== ex5 ======")
	ex5()
}