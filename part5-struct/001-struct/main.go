package main

import "fmt"

// type person
type person struct {
	first string
	last  string
	age   int
}

// type secretAgent
type secretAgent struct {
	person
	ltk bool
}

// 1. struct
// like class, but we say we've created a value of type
func structDemo() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}
	p2 := person{
		first: "miss",
		last:  "moneypenny",
		age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p2.first)
}

// 2. embedded structs
func embeddedStructs() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk) // promoted
}

// 3. anonymous struct
// avoid code pollution
func anonymousStruct() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}
	fmt.Println(p1)

}

func main() {
	structDemo()
	embeddedStructs()
	anonymousStruct()
}
