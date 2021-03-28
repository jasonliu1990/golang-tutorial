package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

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
}

// ex1
//Starting with this code, marshal the []user to JSON.
//There is a little bit of a curve ball in this hands-on exercise
type user struct {
	First string
	Age   int
}

func ex1() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}
	users := []user{u1, u2, u3}

	fmt.Println(users)
	// your code goes here
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

// ex2
// Starting with this code, unmarshal the JSON into a Go data structure.
type user2 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func ex2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs2 := []byte(s)
	var u2 []user2
	err := json.Unmarshal(bs2, &u2)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range u2 {
		fmt.Println("user ", i)
		fmt.Println(v.First, v.Last, v.Age)
		for _, s := range v.Sayings {
			fmt.Println(s, "\t")
		}
	}
}

// ex3
// Starting with this code, encode to JSON the []user
//sending the results to Stdout.
//Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

type user3 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func ex3() {
	u1 := user3{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user3{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	u3 := user3{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user3{u1, u2, u3}
	//fmt.Println(users)
	// your code goes here
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

// ex4
// Starting with this code,
// sort the []int and []string for each person.
func ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

// ex5
//Starting with this code, sort the []user by
//	1. age
//	2. last
//Also sort each []string “Sayings” for each user
//	1. print everything in a way that is pleasant

type user5 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ByAge []user5

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user5

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

func ex5() {
	u1 := user5{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user5{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}
	u3 := user5{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user5{u1, u2, u3}
	//fmt.Println(users)
	//your code goes here
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Println("\t", s)
		}
	}

	fmt.Println("-------------")
	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Println("\t", s)
		}
	}

	fmt.Println("-------------")

	sort.Sort(ByLast(users))
	for _, u := range users {
		fmt.Println(u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, s := range u.Sayings {
			fmt.Println("\t", s)
		}
	}
}
