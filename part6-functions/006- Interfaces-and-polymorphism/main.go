package main

import "fmt"

// keyword identifier type
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak(){
	fmt.Println("I am", s.first, s.last, "- the person secretAgent")
}

func (p person) speak(){
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

type human interface {
	speak()
}

func bar(h human){
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

type hotdog int
// a value can be more than one type
func main() {
	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"miss",
			"moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr,",
		last: "Yes",
	}

	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y  int
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

