package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age int
}

type ByAge []Person
func (a ByAge) Len()int {return len(a)}
func (a ByAge) Swap(i, j int){a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool {return  a[i].Age < a[j].Age }

type ByName []Person
func (bn ByName) Len()int {return len(bn)}
func (bn ByName) Swap(i, j int){bn[i], bn[j] = bn[j], bn[i]}
func (bn ByName) Less(i, j int) bool {return  bn[i].First < bn[j].First }

func main()  {
	xi := []int{4, 5, 2, 1, 99, 42, 16, 52}
	xs := []string{"james", "bond", "m", "dr", "q", "moneypenny"}
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Ints(xi)
	sort.Strings(xs)
	fmt.Println("-----sort-------")
	fmt.Println(xi)
	fmt.Println(xs)


	fmt.Println("-----sort custom-----")
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}
	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)
}


