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

func main() {
	fmt.Println("====== ex1 ======")
	ex1()
	fmt.Println("====== ex2 ======")
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

}