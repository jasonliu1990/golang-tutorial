package main

import "fmt"

// ex1
// Print every number from 1 to 100
func ex1() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

// ex2
// Print every rune code point of the uppercase alphabet three times.
// A: 65; Z:90
func ex2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

// ex3
// for condition { }
// Have it print out the years you have been alive.
func ex3() {
	y := 1990
	for y < 2022 {
		fmt.Println(y)
		y++
	}
}

// ex4
// for { }
// Have it print out the years you have been alive.
func ex4() {
	y := 1990
	for {
		fmt.Println(y)
		y++
		if y == 2022 {
			break
		}
	}
}

// ex5
// Print out the remainder (modulus) which is found for each number
// between 10 and 100 when it is divided by 4.
func ex5() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

// ex6
// Create a program that shows the “if statement” in action.
func ex6() {
	a := 42
	if a == 42 {
		fmt.Println("OK")
	}
}

// ex7
// Building on the previous hands-on exercise, create a program that
// uses “else if” and “else”.
func ex7() {
	a := 42
	if a == 41 {
		fmt.Println("a = 41")
	} else if a == 42 {
		fmt.Println("a = 42")
	} else {
		fmt.Println("a not in (41, 42)")
	}
}

// ex8
// Create a program that uses a switch statement
// with no switch expression specified.
func ex8() {
	x := 42
	switch {
	case x < 42:
		fmt.Println("x < 42")
	case x >= 42:
		fmt.Println("x >= 42")
	}
}

// ex8
//Create a program that uses a switch statement
//with the switch expression specified as a variable of TYPE string
//with the IDENTIFIER “favSport”.
func ex9() {
	favSport := "basketball"
	switch favSport {
	case "basketball":
		fmt.Println("basketball!!!")
	case "baseball":
		fmt.Println("baseball!!!")
	}
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
	fmt.Println("====== ex7 ======")
	ex7()
	fmt.Println("====== ex8 ======")
	ex8()
	fmt.Println("====== ex9 ======")
	ex9()
}
