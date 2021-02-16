package main

import (
	"fmt"
)

// ex1
// Write a program that prints a number in decimal, binary, and hex
func ex1() {
	x := 42
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
}

// ex2
//Using the following operators,
//write expressions and assign their values to variables:
// g. ==
// h. <=
// i. >=
// j. !=
// k. <
// l. >
//Now print each of the variables.
func ex2() {
	g := (42 == 42)
	h := (43 <= 42)
	i := (43 >= 42)
	j := (43 != 42)
	k := (43 < 42)
	l := (43 > 42)
	println(g, h, i, j, k, l)
}

// ex3
//Create TYPED and UNTYPED constants.
//Print the values of the constants.
const (
	a     = 4
	b int = 42
)

func ex3() {
	println(a, b)
}

// ex4
//Write a program that
//	a. assigns an int to a variable
//  b. prints that int in decimal, binary, and hex
//	c. shifts the bits of that int over 1 position to the left,
//     and assigns that to a variable
//	d. prints that variable in decimal, binary, and hex
func ex4() {
	x := 42
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
	x_shift := x << 1
	fmt.Printf("%d\t%b\t%#X\n", x_shift, x_shift, x_shift)

}

// ex5
//Create a variable of type string using a raw string literal. Print it.
func ex5() {
	x := `raw string 
literal`
	fmt.Println(x)
}

// ex6
//Using iota,
//create 4 constants for the NEXT 4 years. Print the constant values.

const (
	_ = iota
	i = 2021 + iota
	j = 2021 + iota
	k = 2021 + iota
	l = 2021 + iota
)

func ex6() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

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
	fmt.Println("====== ex6 ======")
	ex6()

}
