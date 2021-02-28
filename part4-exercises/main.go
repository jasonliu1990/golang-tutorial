package main

import "fmt"

// ex1
// a. Using a COMPOSITE LITERAL:
// 		1. create an ARRAY which holds 5 VALUES of TYPE int
// 		2. assign VALUES to each index position.
// b. Range over the array and print the values out.
// c. Using format printing
//		1. print out the TYPE of the array

func ex1() {
	x := [5]int{1, 2, 4, 8, 16}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)

}

// ex2
// a. Using a COMPOSITE LITERAL:
//		1. create a SLICE of TYPE int
//		2. assign 10 VALUES
// b. Range over the slice and print the values out.
// c. Using format printing
//		1. print out the TYPE of the slice
func ex2() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

// ex3
// Using the code from the previous example, use SLICING to create the following new slices
// which are then printed:
//	1. [42 43 44 45 46]
//	2. [47 48 49 50 51]
//	3. [44 45 46 47 48]
//	4. [43 44 45 46 47]
func ex3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

// ex4
//Follow these steps:
// a. start with this slice
//		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// b. append to that slice this value
//		52
// c. print out the slice
// d. in ONE STATEMENT append to that slice these values
//		1. 53
//		2. 54
//		3. 55
// e. print out the slice
// f. append to the slice this slice
//		1. y := []int{56, 57, 58, 59, 60}
//		2. print out the slice
func ex4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

// ex5
//To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
//follow these steps:
// a. start with this slice
//		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
// b. use APPEND & SLICING to get these values here which you should ASSIGN to a variable
//    “y” and then print:
//		[42, 43, 44, 48, 49, 50, 51]
func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:4], x[6:]...)
	fmt.Println(y)

}

// ex6
//Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:
func ex6() {
	x := make([]string, 50, 500)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

// ex7
// Create a slice of a slice of string ([][]string). Store the following data in the
// multi-dimensional slice:
//
//		"James", "Bond", "Shaken, not stirred"
//		"Miss", "Moneypenny", "Helloooooo, James."
//
//Range over the records, then range over the data in each record.
func ex7() {
	p1 := []string{"James", "Bond", "Shaken, not stirred"}
	p2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{p1, p2}
	for i, r := range x {
		fmt.Println("record:", i)
		for i, v := range r {
			fmt.Println(i, v)
		}
	}
}

// ex8
// Create a map with a key of TYPE string which is a person’s “last_first” name, and a
// value of TYPE []string which stores their favorite things. Store three records in your
// map. Print out all of the values, along with their index position in the slice.
//
//`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
func ex8() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}

// ex9
// Using the code from the previous example, add a record to your map.
// Now print the map out using the “range” loop
func ex9() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["Jaz_L"] = []string{`python`, `golang`, `js`}
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}

// ex 10
// Using the code from the previous example, delete a record from your map.
// Now print the map out using the “range” loop
func ex10() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	delete(m, `no_dr`)
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
func main() {
	fmt.Println("======ex1======")
	ex1()
	fmt.Println("======ex2======")
	ex2()
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
